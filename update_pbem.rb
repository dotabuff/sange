require 'pp'
require 'open3'
require 'strscan'

enums = Hash.new{|h,k| h[k] = {} }

Dir.glob('dota/**/*.proto') do |filename|
  proto = File.read(filename)
  s = StringScanner.new(proto)
  until s.eos?
    if s.scan(/^enum\s+(\S+)\s+\{/)
      enum = s[1]
      until s.scan(/\}/)
        if s.scan(/^\s*([\S]+)\s*=\s*(\d+)\s*;/)
          enums[enum][s[1]] = s[2]
        else
          s.scan(/./m)
        end
      end
    else
      s.scan(/./m)
    end
  end
end

broken = Hash.new{|h,k| h[k] = {} }
broken["SVC_Messages"] = %w[svc_EntityMessage]
broken["EBaseUserMessages"] = %w[UM_MAX_BASE UM_CloseCaptionDirect]
broken["EDotaUserMessages"] = %w[DOTA_UM_AddUnitToSelection DOTA_UM_CharacterSpeakConcept DOTA_UM_GamerulesStateChanged DOTA_UM_TournamentDrop]
broken["EDemoCommands"] = %w[DEM_SignonPacket DEM_Max DEM_IsCompressed]

map_type_options = {
  DEM: {enum: 'EDemoCommands',     iprefix: 'CDemo',         tprefix: 'DEM_'},
  NET: {enum: 'NET_Messages',      iprefix: 'CNETMsg_',      tprefix: 'NET_'},
  SVC: {enum: 'SVC_Messages',      iprefix: 'CSVCMsg_',      tprefix: 'SVC_'},
  BUM: {enum: 'EBaseUserMessages', iprefix: 'CUserMessage',  tprefix: 'UM_'},
  DUM: {enum: 'EDotaUserMessages', iprefix: 'CDOTAUserMsg_', tprefix: 'DOTA_UM_'},
}

map_type = ->(kind, name, value){
  suffix = name.split('_').last
  o = map_type_options[kind]
  if broken[o[:enum]].include?(name)
    warn "Ignoring #{{kind => {name => value}}.inspect}"
    next
  end
  i = [o[:iprefix], suffix].compact.join
  %(&dota.#{i}{})
}

File.open 'core/parser/parser_base_event.go', 'w+' do |io|
  Open3.popen3 'gofmt' do |si, so, se|
    tee = ->(*str){
      puts(*str)
      si.puts(*str)
    }

    pcase = ->(long, short){
      enums[long.to_s].each{|k,v| (s = map_type.(short.to_sym, k, v)) && tee.('case %p: return %s, nil' % [k, s]) }
    }

    dcase = ->(long, short){
      enums[long.to_s].each{|k,v| (s = map_type.(short.to_sym, k, v)) && tee.('case %d: return %s, nil' % [v, s]) }
    }

    tee.(<<-EOS)
// NOTE: This file is generated by tool/update_pbem.rb
package parser

import (
	"code.google.com/p/gogoprotobuf/proto"
	"github.com/dotabuff/sange/dota"
)
    EOS

    tee.('func (p *Parser) AsBaseEvent(commandName string) (proto.Message, error) {')
	  tee.('switch commandName {')
    tee.('case "DEM_SignonPacket": return &SignonPacket{}, nil')
    pcase.('EDemoCommands', :DEM)
    pcase.('NET_Messages', :NET)
    pcase.('SVC_Messages', :SVC)
    pcase.('EBaseUserMessages', :BUM)
    pcase.('EDotaUserMessages', :DUM)
    tee.('}')
    tee.('return nil, Error("Type not found: " + commandName)')
    tee.('}')

    tee.('func (p *Parser) AsBaseEventNETSVC(value int) (proto.Message, error) {')
    tee.('switch value {')
    dcase.('NET_Messages', :NET)
    dcase.('SVC_Messages', :SVC)
    tee.('}')
    tee.('return nil, Error("not found")')
    tee.('}')

    tee.('func (p *Parser) AsBaseEventBUMDUM(value int) (proto.Message, error) {')
    tee.('switch value {')
    dcase.('EBaseUserMessages', :BUM)
    dcase.('EDotaUserMessages', :DUM)
    tee.('}')
    tee.('return nil, Error("Unknown BUMDUM")')
    tee.('}')

    si.close
    io.write(so.read)
    warn se.read
  end
end
