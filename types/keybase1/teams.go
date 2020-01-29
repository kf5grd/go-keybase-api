// Auto-generated to Go types using avdl-compiler v1.4.6 (https://github.com/keybase/node-avdl-compiler)
//   Input file: ../client/protocol/avdl/keybase1/teams.avdl

package keybase1

import (
	"errors"
	"fmt"
)

type TeamRole int

const (
	TeamRole_NONE          TeamRole = 0
	TeamRole_READER        TeamRole = 1
	TeamRole_WRITER        TeamRole = 2
	TeamRole_ADMIN         TeamRole = 3
	TeamRole_OWNER         TeamRole = 4
	TeamRole_BOT           TeamRole = 5
	TeamRole_RESTRICTEDBOT TeamRole = 6
)

func (o TeamRole) DeepCopy() TeamRole { return o }

var TeamRoleMap = map[string]TeamRole{
	"NONE":          0,
	"READER":        1,
	"WRITER":        2,
	"ADMIN":         3,
	"OWNER":         4,
	"BOT":           5,
	"RESTRICTEDBOT": 6,
}

var TeamRoleRevMap = map[TeamRole]string{
	0: "NONE",
	1: "READER",
	2: "WRITER",
	3: "ADMIN",
	4: "OWNER",
	5: "BOT",
	6: "RESTRICTEDBOT",
}

func (e TeamRole) String() string {
	if v, ok := TeamRoleRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamApplication int

const (
	TeamApplication_KBFS                TeamApplication = 1
	TeamApplication_CHAT                TeamApplication = 2
	TeamApplication_SALTPACK            TeamApplication = 3
	TeamApplication_GIT_METADATA        TeamApplication = 4
	TeamApplication_SEITAN_INVITE_TOKEN TeamApplication = 5
	TeamApplication_STELLAR_RELAY       TeamApplication = 6
	TeamApplication_KVSTORE             TeamApplication = 7
)

func (o TeamApplication) DeepCopy() TeamApplication { return o }

var TeamApplicationMap = map[string]TeamApplication{
	"KBFS":                1,
	"CHAT":                2,
	"SALTPACK":            3,
	"GIT_METADATA":        4,
	"SEITAN_INVITE_TOKEN": 5,
	"STELLAR_RELAY":       6,
	"KVSTORE":             7,
}

var TeamApplicationRevMap = map[TeamApplication]string{
	1: "KBFS",
	2: "CHAT",
	3: "SALTPACK",
	4: "GIT_METADATA",
	5: "SEITAN_INVITE_TOKEN",
	6: "STELLAR_RELAY",
	7: "KVSTORE",
}

func (e TeamApplication) String() string {
	if v, ok := TeamApplicationRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamStatus int

const (
	TeamStatus_NONE      TeamStatus = 0
	TeamStatus_LIVE      TeamStatus = 1
	TeamStatus_DELETED   TeamStatus = 2
	TeamStatus_ABANDONED TeamStatus = 3
)

func (o TeamStatus) DeepCopy() TeamStatus { return o }

var TeamStatusMap = map[string]TeamStatus{
	"NONE":      0,
	"LIVE":      1,
	"DELETED":   2,
	"ABANDONED": 3,
}

var TeamStatusRevMap = map[TeamStatus]string{
	0: "NONE",
	1: "LIVE",
	2: "DELETED",
	3: "ABANDONED",
}

func (e TeamStatus) String() string {
	if v, ok := TeamStatusRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type AuditMode int

const (
	AuditMode_STANDARD           AuditMode = 0
	AuditMode_JUST_CREATED       AuditMode = 1
	AuditMode_SKIP               AuditMode = 2
	AuditMode_STANDARD_NO_HIDDEN AuditMode = 3
)

func (o AuditMode) DeepCopy() AuditMode { return o }

var AuditModeMap = map[string]AuditMode{
	"STANDARD":           0,
	"JUST_CREATED":       1,
	"SKIP":               2,
	"STANDARD_NO_HIDDEN": 3,
}

var AuditModeRevMap = map[AuditMode]string{
	0: "STANDARD",
	1: "JUST_CREATED",
	2: "SKIP",
	3: "STANDARD_NO_HIDDEN",
}

func (e AuditMode) String() string {
	if v, ok := AuditModeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type PerTeamKeyGeneration int

func (o PerTeamKeyGeneration) DeepCopy() PerTeamKeyGeneration {
	return o
}

type PTKType int

const (
	PTKType_READER PTKType = 0
)

func (o PTKType) DeepCopy() PTKType { return o }

var PTKTypeMap = map[string]PTKType{
	"READER": 0,
}

var PTKTypeRevMap = map[PTKType]string{
	0: "READER",
}

func (e PTKType) String() string {
	if v, ok := PTKTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type PerTeamSeedCheckVersion int

const (
	PerTeamSeedCheckVersion_V1 PerTeamSeedCheckVersion = 1
)

func (o PerTeamSeedCheckVersion) DeepCopy() PerTeamSeedCheckVersion { return o }

var PerTeamSeedCheckVersionMap = map[string]PerTeamSeedCheckVersion{
	"V1": 1,
}

var PerTeamSeedCheckVersionRevMap = map[PerTeamSeedCheckVersion]string{
	1: "V1",
}

func (e PerTeamSeedCheckVersion) String() string {
	if v, ok := PerTeamSeedCheckVersionRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type PerTeamSeedCheck struct {
	Version PerTeamSeedCheckVersion `codec:"version" json:"version"`
	Value   PerTeamSeedCheckValue   `codec:"value" json:"value"`
}

func (o PerTeamSeedCheck) DeepCopy() PerTeamSeedCheck {
	return PerTeamSeedCheck{
		Version: o.Version.DeepCopy(),
		Value:   o.Value.DeepCopy(),
	}
}

type PerTeamSeedCheckValue []byte

func (o PerTeamSeedCheckValue) DeepCopy() PerTeamSeedCheckValue {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type PerTeamSeedCheckValuePostImage []byte

func (o PerTeamSeedCheckValuePostImage) DeepCopy() PerTeamSeedCheckValuePostImage {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type PerTeamSeedCheckPostImage struct {
	Value   PerTeamSeedCheckValuePostImage `codec:"h" json:"h"`
	Version PerTeamSeedCheckVersion        `codec:"v" json:"v"`
}

func (o PerTeamSeedCheckPostImage) DeepCopy() PerTeamSeedCheckPostImage {
	return PerTeamSeedCheckPostImage{
		Value:   o.Value.DeepCopy(),
		Version: o.Version.DeepCopy(),
	}
}

type TeamApplicationKey struct {
	Application   TeamApplication      `codec:"application" json:"application"`
	KeyGeneration PerTeamKeyGeneration `codec:"keyGeneration" json:"keyGeneration"`
	Key           Bytes32              `codec:"key" json:"key"`
}

func (o TeamApplicationKey) DeepCopy() TeamApplicationKey {
	return TeamApplicationKey{
		Application:   o.Application.DeepCopy(),
		KeyGeneration: o.KeyGeneration.DeepCopy(),
		Key:           o.Key.DeepCopy(),
	}
}

type MaskB64 []byte

func (o MaskB64) DeepCopy() MaskB64 {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type TeamInviteID string

func (o TeamInviteID) DeepCopy() TeamInviteID {
	return o
}

type ReaderKeyMask struct {
	Application TeamApplication      `codec:"application" json:"application"`
	Generation  PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Mask        MaskB64              `codec:"mask" json:"mask"`
}

func (o ReaderKeyMask) DeepCopy() ReaderKeyMask {
	return ReaderKeyMask{
		Application: o.Application.DeepCopy(),
		Generation:  o.Generation.DeepCopy(),
		Mask:        o.Mask.DeepCopy(),
	}
}

type PerTeamKey struct {
	Gen    PerTeamKeyGeneration `codec:"gen" json:"gen"`
	Seqno  Seqno                `codec:"seqno" json:"seqno"`
	SigKID KID                  `codec:"sigKID" json:"sigKID"`
	EncKID KID                  `codec:"encKID" json:"encKID"`
}

func (o PerTeamKey) DeepCopy() PerTeamKey {
	return PerTeamKey{
		Gen:    o.Gen.DeepCopy(),
		Seqno:  o.Seqno.DeepCopy(),
		SigKID: o.SigKID.DeepCopy(),
		EncKID: o.EncKID.DeepCopy(),
	}
}

type PerTeamKeyAndCheck struct {
	Ptk   PerTeamKey                `codec:"ptk" json:"ptk"`
	Check PerTeamSeedCheckPostImage `codec:"check" json:"check"`
}

func (o PerTeamKeyAndCheck) DeepCopy() PerTeamKeyAndCheck {
	return PerTeamKeyAndCheck{
		Ptk:   o.Ptk.DeepCopy(),
		Check: o.Check.DeepCopy(),
	}
}

type PerTeamKeySeed [32]byte

func (o PerTeamKeySeed) DeepCopy() PerTeamKeySeed {
	var ret PerTeamKeySeed
	copy(ret[:], o[:])
	return ret
}

type PerTeamKeySeedItem struct {
	Seed       PerTeamKeySeed       `codec:"seed" json:"seed"`
	Generation PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Seqno      Seqno                `codec:"seqno" json:"seqno"`
	Check      *PerTeamSeedCheck    `codec:"check,omitempty" json:"check,omitempty"`
}

func (o PerTeamKeySeedItem) DeepCopy() PerTeamKeySeedItem {
	return PerTeamKeySeedItem{
		Seed:       o.Seed.DeepCopy(),
		Generation: o.Generation.DeepCopy(),
		Seqno:      o.Seqno.DeepCopy(),
		Check: (func(x *PerTeamSeedCheck) *PerTeamSeedCheck {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Check),
	}
}

type TeamMember struct {
	Uid         UID              `codec:"uid" json:"uid"`
	Role        TeamRole         `codec:"role" json:"role"`
	EldestSeqno Seqno            `codec:"eldestSeqno" json:"eldestSeqno"`
	Status      TeamMemberStatus `codec:"status" json:"status"`
	BotSettings *TeamBotSettings `codec:"botSettings,omitempty" json:"botSettings,omitempty"`
}

func (o TeamMember) DeepCopy() TeamMember {
	return TeamMember{
		Uid:         o.Uid.DeepCopy(),
		Role:        o.Role.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		Status:      o.Status.DeepCopy(),
		BotSettings: (func(x *TeamBotSettings) *TeamBotSettings {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.BotSettings),
	}
}

type TeamMembers struct {
	Owners         []UserVersion `codec:"owners" json:"owners"`
	Admins         []UserVersion `codec:"admins" json:"admins"`
	Writers        []UserVersion `codec:"writers" json:"writers"`
	Readers        []UserVersion `codec:"readers" json:"readers"`
	Bots           []UserVersion `codec:"bots" json:"bots"`
	RestrictedBots []UserVersion `codec:"restrictedBots" json:"restrictedBots"`
}

func (o TeamMembers) DeepCopy() TeamMembers {
	return TeamMembers{
		Owners: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Owners),
		Admins: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Admins),
		Writers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		Readers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Readers),
		Bots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Bots),
		RestrictedBots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.RestrictedBots),
	}
}

type TeamMemberStatus int

const (
	TeamMemberStatus_ACTIVE  TeamMemberStatus = 0
	TeamMemberStatus_RESET   TeamMemberStatus = 1
	TeamMemberStatus_DELETED TeamMemberStatus = 2
)

func (o TeamMemberStatus) DeepCopy() TeamMemberStatus { return o }

var TeamMemberStatusMap = map[string]TeamMemberStatus{
	"ACTIVE":  0,
	"RESET":   1,
	"DELETED": 2,
}

var TeamMemberStatusRevMap = map[TeamMemberStatus]string{
	0: "ACTIVE",
	1: "RESET",
	2: "DELETED",
}

func (e TeamMemberStatus) String() string {
	if v, ok := TeamMemberStatusRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamMemberDetails struct {
	Uv       UserVersion      `codec:"uv" json:"uv"`
	Username string           `codec:"username" json:"username"`
	FullName FullName         `codec:"fullName" json:"fullName"`
	NeedsPUK bool             `codec:"needsPUK" json:"needsPUK"`
	Status   TeamMemberStatus `codec:"status" json:"status"`
}

func (o TeamMemberDetails) DeepCopy() TeamMemberDetails {
	return TeamMemberDetails{
		Uv:       o.Uv.DeepCopy(),
		Username: o.Username,
		FullName: o.FullName.DeepCopy(),
		NeedsPUK: o.NeedsPUK,
		Status:   o.Status.DeepCopy(),
	}
}

type TeamMembersDetails struct {
	Owners         []TeamMemberDetails `codec:"owners" json:"owners"`
	Admins         []TeamMemberDetails `codec:"admins" json:"admins"`
	Writers        []TeamMemberDetails `codec:"writers" json:"writers"`
	Readers        []TeamMemberDetails `codec:"readers" json:"readers"`
	Bots           []TeamMemberDetails `codec:"bots" json:"bots"`
	RestrictedBots []TeamMemberDetails `codec:"restrictedBots" json:"restrictedBots"`
}

func (o TeamMembersDetails) DeepCopy() TeamMembersDetails {
	return TeamMembersDetails{
		Owners: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Owners),
		Admins: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Admins),
		Writers: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		Readers: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Readers),
		Bots: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Bots),
		RestrictedBots: (func(x []TeamMemberDetails) []TeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]TeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.RestrictedBots),
	}
}

type TeamDetails struct {
	Name                   string                               `codec:"name" json:"name"`
	Members                TeamMembersDetails                   `codec:"members" json:"members"`
	KeyGeneration          PerTeamKeyGeneration                 `codec:"keyGeneration" json:"keyGeneration"`
	AnnotatedActiveInvites map[TeamInviteID]AnnotatedTeamInvite `codec:"annotatedActiveInvites" json:"annotatedActiveInvites"`
	Settings               TeamSettings                         `codec:"settings" json:"settings"`
	Showcase               TeamShowcase                         `codec:"showcase" json:"showcase"`
}

func (o TeamDetails) DeepCopy() TeamDetails {
	return TeamDetails{
		Name:          o.Name,
		Members:       o.Members.DeepCopy(),
		KeyGeneration: o.KeyGeneration.DeepCopy(),
		AnnotatedActiveInvites: (func(x map[TeamInviteID]AnnotatedTeamInvite) map[TeamInviteID]AnnotatedTeamInvite {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]AnnotatedTeamInvite, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.AnnotatedActiveInvites),
		Settings: o.Settings.DeepCopy(),
		Showcase: o.Showcase.DeepCopy(),
	}
}

type UserVersionPercentForm string

func (o UserVersionPercentForm) DeepCopy() UserVersionPercentForm {
	return o
}

type TeamChangeReq struct {
	Owners           []UserVersion                           `codec:"owners" json:"owners"`
	Admins           []UserVersion                           `codec:"admins" json:"admins"`
	Writers          []UserVersion                           `codec:"writers" json:"writers"`
	Readers          []UserVersion                           `codec:"readers" json:"readers"`
	Bots             []UserVersion                           `codec:"bots" json:"bots"`
	RestrictedBots   map[UserVersion]TeamBotSettings         `codec:"restrictedBots" json:"restrictedBots"`
	None             []UserVersion                           `codec:"none" json:"none"`
	CompletedInvites map[TeamInviteID]UserVersionPercentForm `codec:"completedInvites" json:"completedInvites"`
}

func (o TeamChangeReq) DeepCopy() TeamChangeReq {
	return TeamChangeReq{
		Owners: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Owners),
		Admins: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Admins),
		Writers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		Readers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Readers),
		Bots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Bots),
		RestrictedBots: (func(x map[UserVersion]TeamBotSettings) map[UserVersion]TeamBotSettings {
			if x == nil {
				return nil
			}
			ret := make(map[UserVersion]TeamBotSettings, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.RestrictedBots),
		None: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.None),
		CompletedInvites: (func(x map[TeamInviteID]UserVersionPercentForm) map[TeamInviteID]UserVersionPercentForm {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]UserVersionPercentForm, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.CompletedInvites),
	}
}

type TeamPlusApplicationKeys struct {
	Id                 TeamID               `codec:"id" json:"id"`
	Name               string               `codec:"name" json:"name"`
	Implicit           bool                 `codec:"implicit" json:"implicit"`
	Public             bool                 `codec:"public" json:"public"`
	Application        TeamApplication      `codec:"application" json:"application"`
	Writers            []UserVersion        `codec:"writers" json:"writers"`
	OnlyReaders        []UserVersion        `codec:"onlyReaders" json:"onlyReaders"`
	OnlyRestrictedBots []UserVersion        `codec:"onlyRestrictedBots" json:"onlyRestrictedBots"`
	ApplicationKeys    []TeamApplicationKey `codec:"applicationKeys" json:"applicationKeys"`
}

func (o TeamPlusApplicationKeys) DeepCopy() TeamPlusApplicationKeys {
	return TeamPlusApplicationKeys{
		Id:          o.Id.DeepCopy(),
		Name:        o.Name,
		Implicit:    o.Implicit,
		Public:      o.Public,
		Application: o.Application.DeepCopy(),
		Writers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Writers),
		OnlyReaders: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.OnlyReaders),
		OnlyRestrictedBots: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.OnlyRestrictedBots),
		ApplicationKeys: (func(x []TeamApplicationKey) []TeamApplicationKey {
			if x == nil {
				return nil
			}
			ret := make([]TeamApplicationKey, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ApplicationKeys),
	}
}

type TeamData struct {
	Subversion                int                                                  `codec:"v" json:"v"`
	Frozen                    bool                                                 `codec:"frozen" json:"frozen"`
	Tombstoned                bool                                                 `codec:"tombstoned" json:"tombstoned"`
	Secretless                bool                                                 `codec:"secretless" json:"secretless"`
	Name                      TeamName                                             `codec:"name" json:"name"`
	Chain                     TeamSigChainState                                    `codec:"chain" json:"chain"`
	PerTeamKeySeedsUnverified map[PerTeamKeyGeneration]PerTeamKeySeedItem          `codec:"perTeamKeySeeds" json:"perTeamKeySeedsUnverified"`
	ReaderKeyMasks            map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 `codec:"readerKeyMasks" json:"readerKeyMasks"`
	LatestSeqnoHint           Seqno                                                `codec:"latestSeqnoHint" json:"latestSeqnoHint"`
	CachedAt                  Time                                                 `codec:"cachedAt" json:"cachedAt"`
	TlfCryptKeys              map[TeamApplication][]CryptKey                       `codec:"tlfCryptKeys" json:"tlfCryptKeys"`
}

func (o TeamData) DeepCopy() TeamData {
	return TeamData{
		Subversion: o.Subversion,
		Frozen:     o.Frozen,
		Tombstoned: o.Tombstoned,
		Secretless: o.Secretless,
		Name:       o.Name.DeepCopy(),
		Chain:      o.Chain.DeepCopy(),
		PerTeamKeySeedsUnverified: (func(x map[PerTeamKeyGeneration]PerTeamKeySeedItem) map[PerTeamKeyGeneration]PerTeamKeySeedItem {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKeySeedItem, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeySeedsUnverified),
		ReaderKeyMasks: (func(x map[TeamApplication]map[PerTeamKeyGeneration]MaskB64) map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication]map[PerTeamKeyGeneration]MaskB64, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x map[PerTeamKeyGeneration]MaskB64) map[PerTeamKeyGeneration]MaskB64 {
					if x == nil {
						return nil
					}
					ret := make(map[PerTeamKeyGeneration]MaskB64, len(x))
					for k, v := range x {
						kCopy := k.DeepCopy()
						vCopy := v.DeepCopy()
						ret[kCopy] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ReaderKeyMasks),
		LatestSeqnoHint: o.LatestSeqnoHint.DeepCopy(),
		CachedAt:        o.CachedAt.DeepCopy(),
		TlfCryptKeys: (func(x map[TeamApplication][]CryptKey) map[TeamApplication][]CryptKey {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication][]CryptKey, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []CryptKey) []CryptKey {
					if x == nil {
						return nil
					}
					ret := make([]CryptKey, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.TlfCryptKeys),
	}
}

type FastTeamData struct {
	Frozen                     bool                                                 `codec:"frozen" json:"frozen"`
	Subversion                 int                                                  `codec:"subversion" json:"subversion"`
	Tombstoned                 bool                                                 `codec:"tombstoned" json:"tombstoned"`
	Name                       TeamName                                             `codec:"name" json:"name"`
	Chain                      FastTeamSigChainState                                `codec:"chain" json:"chain"`
	PerTeamKeySeedsUnverified  map[PerTeamKeyGeneration]PerTeamKeySeed              `codec:"perTeamKeySeeds" json:"perTeamKeySeedsUnverified"`
	MaxContinuousPTKGeneration PerTeamKeyGeneration                                 `codec:"maxContinuousPTKGeneration" json:"maxContinuousPTKGeneration"`
	SeedChecks                 map[PerTeamKeyGeneration]PerTeamSeedCheck            `codec:"seedChecks" json:"seedChecks"`
	LatestKeyGeneration        PerTeamKeyGeneration                                 `codec:"latestKeyGeneration" json:"latestKeyGeneration"`
	ReaderKeyMasks             map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 `codec:"readerKeyMasks" json:"readerKeyMasks"`
	LatestSeqnoHint            Seqno                                                `codec:"latestSeqnoHint" json:"latestSeqnoHint"`
	CachedAt                   Time                                                 `codec:"cachedAt" json:"cachedAt"`
	LoadedLatest               bool                                                 `codec:"loadedLatest" json:"loadedLatest"`
}

func (o FastTeamData) DeepCopy() FastTeamData {
	return FastTeamData{
		Frozen:     o.Frozen,
		Subversion: o.Subversion,
		Tombstoned: o.Tombstoned,
		Name:       o.Name.DeepCopy(),
		Chain:      o.Chain.DeepCopy(),
		PerTeamKeySeedsUnverified: (func(x map[PerTeamKeyGeneration]PerTeamKeySeed) map[PerTeamKeyGeneration]PerTeamKeySeed {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKeySeed, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeySeedsUnverified),
		MaxContinuousPTKGeneration: o.MaxContinuousPTKGeneration.DeepCopy(),
		SeedChecks: (func(x map[PerTeamKeyGeneration]PerTeamSeedCheck) map[PerTeamKeyGeneration]PerTeamSeedCheck {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamSeedCheck, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.SeedChecks),
		LatestKeyGeneration: o.LatestKeyGeneration.DeepCopy(),
		ReaderKeyMasks: (func(x map[TeamApplication]map[PerTeamKeyGeneration]MaskB64) map[TeamApplication]map[PerTeamKeyGeneration]MaskB64 {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication]map[PerTeamKeyGeneration]MaskB64, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x map[PerTeamKeyGeneration]MaskB64) map[PerTeamKeyGeneration]MaskB64 {
					if x == nil {
						return nil
					}
					ret := make(map[PerTeamKeyGeneration]MaskB64, len(x))
					for k, v := range x {
						kCopy := k.DeepCopy()
						vCopy := v.DeepCopy()
						ret[kCopy] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ReaderKeyMasks),
		LatestSeqnoHint: o.LatestSeqnoHint.DeepCopy(),
		CachedAt:        o.CachedAt.DeepCopy(),
		LoadedLatest:    o.LoadedLatest,
	}
}

type RatchetType int

const (
	RatchetType_MAIN        RatchetType = 0
	RatchetType_BLINDED     RatchetType = 1
	RatchetType_SELF        RatchetType = 2
	RatchetType_UNCOMMITTED RatchetType = 3
)

func (o RatchetType) DeepCopy() RatchetType { return o }

var RatchetTypeMap = map[string]RatchetType{
	"MAIN":        0,
	"BLINDED":     1,
	"SELF":        2,
	"UNCOMMITTED": 3,
}

var RatchetTypeRevMap = map[RatchetType]string{
	0: "MAIN",
	1: "BLINDED",
	2: "SELF",
	3: "UNCOMMITTED",
}

func (e RatchetType) String() string {
	if v, ok := RatchetTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type HiddenTeamChainRatchetSet struct {
	Ratchets map[RatchetType]LinkTripleAndTime `codec:"ratchets" json:"ratchets"`
}

func (o HiddenTeamChainRatchetSet) DeepCopy() HiddenTeamChainRatchetSet {
	return HiddenTeamChainRatchetSet{
		Ratchets: (func(x map[RatchetType]LinkTripleAndTime) map[RatchetType]LinkTripleAndTime {
			if x == nil {
				return nil
			}
			ret := make(map[RatchetType]LinkTripleAndTime, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Ratchets),
	}
}

type HiddenTeamChain struct {
	Id                 TeamID                         `codec:"id" json:"id"`
	Subversion         int                            `codec:"subversion" json:"subversion"`
	Public             bool                           `codec:"public" json:"public"`
	Frozen             bool                           `codec:"frozen" json:"frozen"`
	Tombstoned         bool                           `codec:"tombstoned" json:"tombstoned"`
	Last               Seqno                          `codec:"last" json:"last"`
	LastFull           Seqno                          `codec:"lastFull" json:"lastFull"`
	LatestSeqnoHint    Seqno                          `codec:"latestSeqnoHint" json:"latestSeqnoHint"`
	LastCommittedSeqno Seqno                          `codec:"lastCommittedSeqno" json:"lastCommittedSeqno"`
	LinkReceiptTimes   map[Seqno]Time                 `codec:"linkReceiptTimes" json:"linkReceiptTimes"`
	LastPerTeamKeys    map[PTKType]Seqno              `codec:"lastPerTeamKeys" json:"lastPerTeamKeys"`
	Outer              map[Seqno]LinkID               `codec:"outer" json:"outer"`
	Inner              map[Seqno]HiddenTeamChainLink  `codec:"inner" json:"inner"`
	ReaderPerTeamKeys  map[PerTeamKeyGeneration]Seqno `codec:"readerPerTeamKeys" json:"readerPerTeamKeys"`
	RatchetSet         HiddenTeamChainRatchetSet      `codec:"ratchetSet" json:"ratchetSet"`
	CachedAt           Time                           `codec:"cachedAt" json:"cachedAt"`
	NeedRotate         bool                           `codec:"needRotate" json:"needRotate"`
	MerkleRoots        map[Seqno]MerkleRootV2         `codec:"merkleRoots" json:"merkleRoots"`
}

func (o HiddenTeamChain) DeepCopy() HiddenTeamChain {
	return HiddenTeamChain{
		Id:                 o.Id.DeepCopy(),
		Subversion:         o.Subversion,
		Public:             o.Public,
		Frozen:             o.Frozen,
		Tombstoned:         o.Tombstoned,
		Last:               o.Last.DeepCopy(),
		LastFull:           o.LastFull.DeepCopy(),
		LatestSeqnoHint:    o.LatestSeqnoHint.DeepCopy(),
		LastCommittedSeqno: o.LastCommittedSeqno.DeepCopy(),
		LinkReceiptTimes: (func(x map[Seqno]Time) map[Seqno]Time {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]Time, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LinkReceiptTimes),
		LastPerTeamKeys: (func(x map[PTKType]Seqno) map[PTKType]Seqno {
			if x == nil {
				return nil
			}
			ret := make(map[PTKType]Seqno, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LastPerTeamKeys),
		Outer: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Outer),
		Inner: (func(x map[Seqno]HiddenTeamChainLink) map[Seqno]HiddenTeamChainLink {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]HiddenTeamChainLink, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Inner),
		ReaderPerTeamKeys: (func(x map[PerTeamKeyGeneration]Seqno) map[PerTeamKeyGeneration]Seqno {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]Seqno, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ReaderPerTeamKeys),
		RatchetSet: o.RatchetSet.DeepCopy(),
		CachedAt:   o.CachedAt.DeepCopy(),
		NeedRotate: o.NeedRotate,
		MerkleRoots: (func(x map[Seqno]MerkleRootV2) map[Seqno]MerkleRootV2 {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]MerkleRootV2, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.MerkleRoots),
	}
}

type LinkTriple struct {
	Seqno   Seqno   `codec:"seqno" json:"seqno"`
	SeqType SeqType `codec:"seqType" json:"seqType"`
	LinkID  LinkID  `codec:"linkID" json:"linkID"`
}

func (o LinkTriple) DeepCopy() LinkTriple {
	return LinkTriple{
		Seqno:   o.Seqno.DeepCopy(),
		SeqType: o.SeqType.DeepCopy(),
		LinkID:  o.LinkID.DeepCopy(),
	}
}

type LinkTripleAndTime struct {
	Triple LinkTriple `codec:"triple" json:"triple"`
	Time   Time       `codec:"time" json:"time"`
}

func (o LinkTripleAndTime) DeepCopy() LinkTripleAndTime {
	return LinkTripleAndTime{
		Triple: o.Triple.DeepCopy(),
		Time:   o.Time.DeepCopy(),
	}
}

type UpPointer struct {
	OurSeqno    Seqno  `codec:"ourSeqno" json:"ourSeqno"`
	ParentID    TeamID `codec:"parentID" json:"parentID"`
	ParentSeqno Seqno  `codec:"parentSeqno" json:"parentSeqno"`
	Deletion    bool   `codec:"deletion" json:"deletion"`
}

func (o UpPointer) DeepCopy() UpPointer {
	return UpPointer{
		OurSeqno:    o.OurSeqno.DeepCopy(),
		ParentID:    o.ParentID.DeepCopy(),
		ParentSeqno: o.ParentSeqno.DeepCopy(),
		Deletion:    o.Deletion,
	}
}

type DownPointer struct {
	Id            TeamID `codec:"id" json:"id"`
	NameComponent string `codec:"nameComponent" json:"nameComponent"`
	IsDeleted     bool   `codec:"isDeleted" json:"isDeleted"`
}

func (o DownPointer) DeepCopy() DownPointer {
	return DownPointer{
		Id:            o.Id.DeepCopy(),
		NameComponent: o.NameComponent,
		IsDeleted:     o.IsDeleted,
	}
}

type Signer struct {
	E Seqno `codec:"e" json:"e"`
	K KID   `codec:"k" json:"k"`
	U UID   `codec:"u" json:"u"`
}

func (o Signer) DeepCopy() Signer {
	return Signer{
		E: o.E.DeepCopy(),
		K: o.K.DeepCopy(),
		U: o.U.DeepCopy(),
	}
}

type HiddenTeamChainLink struct {
	MerkleRoot  MerkleRootV2                   `codec:"m" json:"m"`
	ParentChain LinkTriple                     `codec:"p" json:"p"`
	Signer      Signer                         `codec:"s" json:"s"`
	Ptk         map[PTKType]PerTeamKeyAndCheck `codec:"k" json:"k"`
}

func (o HiddenTeamChainLink) DeepCopy() HiddenTeamChainLink {
	return HiddenTeamChainLink{
		MerkleRoot:  o.MerkleRoot.DeepCopy(),
		ParentChain: o.ParentChain.DeepCopy(),
		Signer:      o.Signer.DeepCopy(),
		Ptk: (func(x map[PTKType]PerTeamKeyAndCheck) map[PTKType]PerTeamKeyAndCheck {
			if x == nil {
				return nil
			}
			ret := make(map[PTKType]PerTeamKeyAndCheck, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Ptk),
	}
}

type FastTeamSigChainState struct {
	ID                      TeamID                                  `codec:"ID" json:"ID"`
	Public                  bool                                    `codec:"public" json:"public"`
	RootAncestor            TeamName                                `codec:"rootAncestor" json:"rootAncestor"`
	NameDepth               int                                     `codec:"nameDepth" json:"nameDepth"`
	Last                    *LinkTriple                             `codec:"last,omitempty" json:"last,omitempty"`
	PerTeamKeys             map[PerTeamKeyGeneration]PerTeamKey     `codec:"perTeamKeys" json:"perTeamKeys"`
	PerTeamKeySeedsVerified map[PerTeamKeyGeneration]PerTeamKeySeed `codec:"perTeamKeySeedsVerified" json:"perTeamKeySeedsVerified"`
	DownPointers            map[Seqno]DownPointer                   `codec:"downPointers" json:"downPointers"`
	LastUpPointer           *UpPointer                              `codec:"lastUpPointer,omitempty" json:"lastUpPointer,omitempty"`
	PerTeamKeyCTime         UnixTime                                `codec:"perTeamKeyCTime" json:"perTeamKeyCTime"`
	LinkIDs                 map[Seqno]LinkID                        `codec:"linkIDs" json:"linkIDs"`
	MerkleInfo              map[Seqno]MerkleRootV2                  `codec:"merkleInfo" json:"merkleInfo"`
}

func (o FastTeamSigChainState) DeepCopy() FastTeamSigChainState {
	return FastTeamSigChainState{
		ID:           o.ID.DeepCopy(),
		Public:       o.Public,
		RootAncestor: o.RootAncestor.DeepCopy(),
		NameDepth:    o.NameDepth,
		Last: (func(x *LinkTriple) *LinkTriple {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Last),
		PerTeamKeys: (func(x map[PerTeamKeyGeneration]PerTeamKey) map[PerTeamKeyGeneration]PerTeamKey {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKey, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeys),
		PerTeamKeySeedsVerified: (func(x map[PerTeamKeyGeneration]PerTeamKeySeed) map[PerTeamKeyGeneration]PerTeamKeySeed {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKeySeed, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeySeedsVerified),
		DownPointers: (func(x map[Seqno]DownPointer) map[Seqno]DownPointer {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]DownPointer, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.DownPointers),
		LastUpPointer: (func(x *UpPointer) *UpPointer {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.LastUpPointer),
		PerTeamKeyCTime: o.PerTeamKeyCTime.DeepCopy(),
		LinkIDs: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LinkIDs),
		MerkleInfo: (func(x map[Seqno]MerkleRootV2) map[Seqno]MerkleRootV2 {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]MerkleRootV2, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.MerkleInfo),
	}
}

type Audit struct {
	Time           Time  `codec:"time" json:"time"`
	MaxMerkleSeqno Seqno `codec:"mms" json:"mms"`
	MaxChainSeqno  Seqno `codec:"mcs" json:"mcs"`
	MaxHiddenSeqno Seqno `codec:"mhs" json:"mhs"`
	MaxMerkleProbe Seqno `codec:"mmp" json:"mmp"`
}

func (o Audit) DeepCopy() Audit {
	return Audit{
		Time:           o.Time.DeepCopy(),
		MaxMerkleSeqno: o.MaxMerkleSeqno.DeepCopy(),
		MaxChainSeqno:  o.MaxChainSeqno.DeepCopy(),
		MaxHiddenSeqno: o.MaxHiddenSeqno.DeepCopy(),
		MaxMerkleProbe: o.MaxMerkleProbe.DeepCopy(),
	}
}

type Probe struct {
	Index           int   `codec:"i" json:"i"`
	TeamSeqno       Seqno `codec:"s" json:"t"`
	TeamHiddenSeqno Seqno `codec:"h" json:"h"`
}

func (o Probe) DeepCopy() Probe {
	return Probe{
		Index:           o.Index,
		TeamSeqno:       o.TeamSeqno.DeepCopy(),
		TeamHiddenSeqno: o.TeamHiddenSeqno.DeepCopy(),
	}
}

type AuditVersion int

const (
	AuditVersion_V0 AuditVersion = 0
	AuditVersion_V1 AuditVersion = 1
	AuditVersion_V2 AuditVersion = 2
	AuditVersion_V3 AuditVersion = 3
	AuditVersion_V4 AuditVersion = 4
)

func (o AuditVersion) DeepCopy() AuditVersion { return o }

var AuditVersionMap = map[string]AuditVersion{
	"V0": 0,
	"V1": 1,
	"V2": 2,
	"V3": 3,
	"V4": 4,
}

var AuditVersionRevMap = map[AuditVersion]string{
	0: "V0",
	1: "V1",
	2: "V2",
	3: "V3",
	4: "V4",
}

func (e AuditVersion) String() string {
	if v, ok := AuditVersionRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type AuditHistory struct {
	ID                TeamID           `codec:"ID" json:"ID"`
	Public            bool             `codec:"public" json:"public"`
	PriorMerkleSeqno  Seqno            `codec:"priorMerkleSeqno" json:"priorMerkleSeqno"`
	Version           AuditVersion     `codec:"version" json:"version"`
	Audits            []Audit          `codec:"audits" json:"audits"`
	PreProbes         map[Seqno]Probe  `codec:"preProbes" json:"preProbes"`
	PostProbes        map[Seqno]Probe  `codec:"postProbes" json:"postProbes"`
	Tails             map[Seqno]LinkID `codec:"tails" json:"tails"`
	HiddenTails       map[Seqno]LinkID `codec:"hiddenTails" json:"hiddenTails"`
	PreProbesToRetry  []Seqno          `codec:"preProbesToRetry" json:"preProbesToRetry"`
	PostProbesToRetry []Seqno          `codec:"postProbesToRetry" json:"postProbesToRetry"`
	SkipUntil         Time             `codec:"skipUntil" json:"skipUntil"`
}

func (o AuditHistory) DeepCopy() AuditHistory {
	return AuditHistory{
		ID:               o.ID.DeepCopy(),
		Public:           o.Public,
		PriorMerkleSeqno: o.PriorMerkleSeqno.DeepCopy(),
		Version:          o.Version.DeepCopy(),
		Audits: (func(x []Audit) []Audit {
			if x == nil {
				return nil
			}
			ret := make([]Audit, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Audits),
		PreProbes: (func(x map[Seqno]Probe) map[Seqno]Probe {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]Probe, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PreProbes),
		PostProbes: (func(x map[Seqno]Probe) map[Seqno]Probe {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]Probe, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PostProbes),
		Tails: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Tails),
		HiddenTails: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.HiddenTails),
		PreProbesToRetry: (func(x []Seqno) []Seqno {
			if x == nil {
				return nil
			}
			ret := make([]Seqno, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.PreProbesToRetry),
		PostProbesToRetry: (func(x []Seqno) []Seqno {
			if x == nil {
				return nil
			}
			ret := make([]Seqno, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.PostProbesToRetry),
		SkipUntil: o.SkipUntil.DeepCopy(),
	}
}

type TeamInviteCategory int

const (
	TeamInviteCategory_NONE    TeamInviteCategory = 0
	TeamInviteCategory_UNKNOWN TeamInviteCategory = 1
	TeamInviteCategory_KEYBASE TeamInviteCategory = 2
	TeamInviteCategory_EMAIL   TeamInviteCategory = 3
	TeamInviteCategory_SBS     TeamInviteCategory = 4
	TeamInviteCategory_SEITAN  TeamInviteCategory = 5
	TeamInviteCategory_PHONE   TeamInviteCategory = 6
)

func (o TeamInviteCategory) DeepCopy() TeamInviteCategory { return o }

var TeamInviteCategoryMap = map[string]TeamInviteCategory{
	"NONE":    0,
	"UNKNOWN": 1,
	"KEYBASE": 2,
	"EMAIL":   3,
	"SBS":     4,
	"SEITAN":  5,
	"PHONE":   6,
}

var TeamInviteCategoryRevMap = map[TeamInviteCategory]string{
	0: "NONE",
	1: "UNKNOWN",
	2: "KEYBASE",
	3: "EMAIL",
	4: "SBS",
	5: "SEITAN",
	6: "PHONE",
}

func (e TeamInviteCategory) String() string {
	if v, ok := TeamInviteCategoryRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamInviteType struct {
	C__       TeamInviteCategory       `codec:"c" json:"c"`
	Unknown__ *string                  `codec:"unknown,omitempty" json:"unknown,omitempty"`
	Sbs__     *TeamInviteSocialNetwork `codec:"sbs,omitempty" json:"sbs,omitempty"`
}

func (o *TeamInviteType) C() (ret TeamInviteCategory, err error) {
	switch o.C__ {
	case TeamInviteCategory_UNKNOWN:
		if o.Unknown__ == nil {
			err = errors.New("unexpected nil value for Unknown__")
			return ret, err
		}
	case TeamInviteCategory_SBS:
		if o.Sbs__ == nil {
			err = errors.New("unexpected nil value for Sbs__")
			return ret, err
		}
	}
	return o.C__, nil
}

func (o TeamInviteType) Unknown() (res string) {
	if o.C__ != TeamInviteCategory_UNKNOWN {
		panic("wrong case accessed")
	}
	if o.Unknown__ == nil {
		return
	}
	return *o.Unknown__
}

func (o TeamInviteType) Sbs() (res TeamInviteSocialNetwork) {
	if o.C__ != TeamInviteCategory_SBS {
		panic("wrong case accessed")
	}
	if o.Sbs__ == nil {
		return
	}
	return *o.Sbs__
}

func NewTeamInviteTypeWithUnknown(v string) TeamInviteType {
	return TeamInviteType{
		C__:       TeamInviteCategory_UNKNOWN,
		Unknown__: &v,
	}
}

func NewTeamInviteTypeWithSbs(v TeamInviteSocialNetwork) TeamInviteType {
	return TeamInviteType{
		C__:   TeamInviteCategory_SBS,
		Sbs__: &v,
	}
}

func NewTeamInviteTypeDefault(c TeamInviteCategory) TeamInviteType {
	return TeamInviteType{
		C__: c,
	}
}

func (o TeamInviteType) DeepCopy() TeamInviteType {
	return TeamInviteType{
		C__: o.C__.DeepCopy(),
		Unknown__: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Unknown__),
		Sbs__: (func(x *TeamInviteSocialNetwork) *TeamInviteSocialNetwork {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Sbs__),
	}
}

type TeamInviteSocialNetwork string

func (o TeamInviteSocialNetwork) DeepCopy() TeamInviteSocialNetwork {
	return o
}

type TeamInviteName string

func (o TeamInviteName) DeepCopy() TeamInviteName {
	return o
}

type TeamInvite struct {
	Role    TeamRole       `codec:"role" json:"role"`
	Id      TeamInviteID   `codec:"id" json:"id"`
	Type    TeamInviteType `codec:"type" json:"type"`
	Name    TeamInviteName `codec:"name" json:"name"`
	Inviter UserVersion    `codec:"inviter" json:"inviter"`
}

func (o TeamInvite) DeepCopy() TeamInvite {
	return TeamInvite{
		Role:    o.Role.DeepCopy(),
		Id:      o.Id.DeepCopy(),
		Type:    o.Type.DeepCopy(),
		Name:    o.Name.DeepCopy(),
		Inviter: o.Inviter.DeepCopy(),
	}
}

type AnnotatedTeamInvite struct {
	Role            TeamRole         `codec:"role" json:"role"`
	Id              TeamInviteID     `codec:"id" json:"id"`
	Type            TeamInviteType   `codec:"type" json:"type"`
	Name            TeamInviteName   `codec:"name" json:"name"`
	Uv              UserVersion      `codec:"uv" json:"uv"`
	Inviter         UserVersion      `codec:"inviter" json:"inviter"`
	InviterUsername string           `codec:"inviterUsername" json:"inviterUsername"`
	TeamName        string           `codec:"teamName" json:"teamName"`
	Status          TeamMemberStatus `codec:"status" json:"status"`
}

func (o AnnotatedTeamInvite) DeepCopy() AnnotatedTeamInvite {
	return AnnotatedTeamInvite{
		Role:            o.Role.DeepCopy(),
		Id:              o.Id.DeepCopy(),
		Type:            o.Type.DeepCopy(),
		Name:            o.Name.DeepCopy(),
		Uv:              o.Uv.DeepCopy(),
		Inviter:         o.Inviter.DeepCopy(),
		InviterUsername: o.InviterUsername,
		TeamName:        o.TeamName,
		Status:          o.Status.DeepCopy(),
	}
}

type TeamEncryptedKBFSKeyset struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

func (o TeamEncryptedKBFSKeyset) DeepCopy() TeamEncryptedKBFSKeyset {
	return TeamEncryptedKBFSKeyset{
		V: o.V,
		E: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.E),
		N: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.N),
	}
}

type TeamGetLegacyTLFUpgrade struct {
	EncryptedKeyset  string               `codec:"encryptedKeyset" json:"encrypted_keyset"`
	TeamGeneration   PerTeamKeyGeneration `codec:"teamGeneration" json:"team_generation"`
	LegacyGeneration int                  `codec:"legacyGeneration" json:"legacy_generation"`
	AppType          TeamApplication      `codec:"appType" json:"app_type"`
}

func (o TeamGetLegacyTLFUpgrade) DeepCopy() TeamGetLegacyTLFUpgrade {
	return TeamGetLegacyTLFUpgrade{
		EncryptedKeyset:  o.EncryptedKeyset,
		TeamGeneration:   o.TeamGeneration.DeepCopy(),
		LegacyGeneration: o.LegacyGeneration,
		AppType:          o.AppType.DeepCopy(),
	}
}

type TeamEncryptedKBFSKeysetHash string

func (o TeamEncryptedKBFSKeysetHash) DeepCopy() TeamEncryptedKBFSKeysetHash {
	return o
}

type TeamLegacyTLFUpgradeChainInfo struct {
	KeysetHash       TeamEncryptedKBFSKeysetHash `codec:"keysetHash" json:"keysetHash"`
	TeamGeneration   PerTeamKeyGeneration        `codec:"teamGeneration" json:"teamGeneration"`
	LegacyGeneration int                         `codec:"legacyGeneration" json:"legacyGeneration"`
	AppType          TeamApplication             `codec:"appType" json:"appType"`
}

func (o TeamLegacyTLFUpgradeChainInfo) DeepCopy() TeamLegacyTLFUpgradeChainInfo {
	return TeamLegacyTLFUpgradeChainInfo{
		KeysetHash:       o.KeysetHash.DeepCopy(),
		TeamGeneration:   o.TeamGeneration.DeepCopy(),
		LegacyGeneration: o.LegacyGeneration,
		AppType:          o.AppType.DeepCopy(),
	}
}

type TeamSigChainState struct {
	Reader                  UserVersion                                       `codec:"reader" json:"reader"`
	Id                      TeamID                                            `codec:"id" json:"id"`
	Implicit                bool                                              `codec:"implicit" json:"implicit"`
	Public                  bool                                              `codec:"public" json:"public"`
	RootAncestor            TeamName                                          `codec:"rootAncestor" json:"rootAncestor"`
	NameDepth               int                                               `codec:"nameDepth" json:"nameDepth"`
	NameLog                 []TeamNameLogPoint                                `codec:"nameLog" json:"nameLog"`
	LastSeqno               Seqno                                             `codec:"lastSeqno" json:"lastSeqno"`
	LastLinkID              LinkID                                            `codec:"lastLinkID" json:"lastLinkID"`
	LastHighSeqno           Seqno                                             `codec:"lastHighSeqno" json:"lastHighSeqno"`
	LastHighLinkID          LinkID                                            `codec:"lastHighLinkID" json:"lastHighLinkID"`
	ParentID                *TeamID                                           `codec:"parentID,omitempty" json:"parentID,omitempty"`
	UserLog                 map[UserVersion][]UserLogPoint                    `codec:"userLog" json:"userLog"`
	SubteamLog              map[TeamID][]SubteamLogPoint                      `codec:"subteamLog" json:"subteamLog"`
	PerTeamKeys             map[PerTeamKeyGeneration]PerTeamKey               `codec:"perTeamKeys" json:"perTeamKeys"`
	MaxPerTeamKeyGeneration PerTeamKeyGeneration                              `codec:"maxPerTeamKeyGeneration" json:"maxPerTeamKeyGeneration"`
	PerTeamKeyCTime         UnixTime                                          `codec:"perTeamKeyCTime" json:"perTeamKeyCTime"`
	LinkIDs                 map[Seqno]LinkID                                  `codec:"linkIDs" json:"linkIDs"`
	StubbedLinks            map[Seqno]bool                                    `codec:"stubbedLinks" json:"stubbedLinks"`
	ActiveInvites           map[TeamInviteID]TeamInvite                       `codec:"activeInvites" json:"activeInvites"`
	ObsoleteInvites         map[TeamInviteID]TeamInvite                       `codec:"obsoleteInvites" json:"obsoleteInvites"`
	Open                    bool                                              `codec:"open" json:"open"`
	OpenTeamJoinAs          TeamRole                                          `codec:"openTeamJoinAs" json:"openTeamJoinAs"`
	Bots                    map[UserVersion]TeamBotSettings                   `codec:"bots" json:"bots"`
	TlfIDs                  []TLFID                                           `codec:"tlfIDs" json:"tlfIDs"`
	TlfLegacyUpgrade        map[TeamApplication]TeamLegacyTLFUpgradeChainInfo `codec:"tlfLegacyUpgrade" json:"tlfLegacyUpgrade"`
	HeadMerkle              *MerkleRootV2                                     `codec:"headMerkle,omitempty" json:"headMerkle,omitempty"`
	MerkleRoots             map[Seqno]MerkleRootV2                            `codec:"merkleRoots" json:"merkleRoots"`
}

func (o TeamSigChainState) DeepCopy() TeamSigChainState {
	return TeamSigChainState{
		Reader:       o.Reader.DeepCopy(),
		Id:           o.Id.DeepCopy(),
		Implicit:     o.Implicit,
		Public:       o.Public,
		RootAncestor: o.RootAncestor.DeepCopy(),
		NameDepth:    o.NameDepth,
		NameLog: (func(x []TeamNameLogPoint) []TeamNameLogPoint {
			if x == nil {
				return nil
			}
			ret := make([]TeamNameLogPoint, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.NameLog),
		LastSeqno:      o.LastSeqno.DeepCopy(),
		LastLinkID:     o.LastLinkID.DeepCopy(),
		LastHighSeqno:  o.LastHighSeqno.DeepCopy(),
		LastHighLinkID: o.LastHighLinkID.DeepCopy(),
		ParentID: (func(x *TeamID) *TeamID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ParentID),
		UserLog: (func(x map[UserVersion][]UserLogPoint) map[UserVersion][]UserLogPoint {
			if x == nil {
				return nil
			}
			ret := make(map[UserVersion][]UserLogPoint, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []UserLogPoint) []UserLogPoint {
					if x == nil {
						return nil
					}
					ret := make([]UserLogPoint, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.UserLog),
		SubteamLog: (func(x map[TeamID][]SubteamLogPoint) map[TeamID][]SubteamLogPoint {
			if x == nil {
				return nil
			}
			ret := make(map[TeamID][]SubteamLogPoint, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []SubteamLogPoint) []SubteamLogPoint {
					if x == nil {
						return nil
					}
					ret := make([]SubteamLogPoint, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.SubteamLog),
		PerTeamKeys: (func(x map[PerTeamKeyGeneration]PerTeamKey) map[PerTeamKeyGeneration]PerTeamKey {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration]PerTeamKey, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.PerTeamKeys),
		MaxPerTeamKeyGeneration: o.MaxPerTeamKeyGeneration.DeepCopy(),
		PerTeamKeyCTime:         o.PerTeamKeyCTime.DeepCopy(),
		LinkIDs: (func(x map[Seqno]LinkID) map[Seqno]LinkID {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]LinkID, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.LinkIDs),
		StubbedLinks: (func(x map[Seqno]bool) map[Seqno]bool {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]bool, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.StubbedLinks),
		ActiveInvites: (func(x map[TeamInviteID]TeamInvite) map[TeamInviteID]TeamInvite {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]TeamInvite, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ActiveInvites),
		ObsoleteInvites: (func(x map[TeamInviteID]TeamInvite) map[TeamInviteID]TeamInvite {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]TeamInvite, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.ObsoleteInvites),
		Open:           o.Open,
		OpenTeamJoinAs: o.OpenTeamJoinAs.DeepCopy(),
		Bots: (func(x map[UserVersion]TeamBotSettings) map[UserVersion]TeamBotSettings {
			if x == nil {
				return nil
			}
			ret := make(map[UserVersion]TeamBotSettings, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Bots),
		TlfIDs: (func(x []TLFID) []TLFID {
			if x == nil {
				return nil
			}
			ret := make([]TLFID, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.TlfIDs),
		TlfLegacyUpgrade: (func(x map[TeamApplication]TeamLegacyTLFUpgradeChainInfo) map[TeamApplication]TeamLegacyTLFUpgradeChainInfo {
			if x == nil {
				return nil
			}
			ret := make(map[TeamApplication]TeamLegacyTLFUpgradeChainInfo, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.TlfLegacyUpgrade),
		HeadMerkle: (func(x *MerkleRootV2) *MerkleRootV2 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.HeadMerkle),
		MerkleRoots: (func(x map[Seqno]MerkleRootV2) map[Seqno]MerkleRootV2 {
			if x == nil {
				return nil
			}
			ret := make(map[Seqno]MerkleRootV2, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.MerkleRoots),
	}
}

type BoxSummaryHash string

func (o BoxSummaryHash) DeepCopy() BoxSummaryHash {
	return o
}

type TeamNameLogPoint struct {
	LastPart TeamNamePart `codec:"lastPart" json:"lastPart"`
	Seqno    Seqno        `codec:"seqno" json:"seqno"`
}

func (o TeamNameLogPoint) DeepCopy() TeamNameLogPoint {
	return TeamNameLogPoint{
		LastPart: o.LastPart.DeepCopy(),
		Seqno:    o.Seqno.DeepCopy(),
	}
}

type UserLogPoint struct {
	Role    TeamRole          `codec:"role" json:"role"`
	SigMeta SignatureMetadata `codec:"sigMeta" json:"sigMeta"`
}

func (o UserLogPoint) DeepCopy() UserLogPoint {
	return UserLogPoint{
		Role:    o.Role.DeepCopy(),
		SigMeta: o.SigMeta.DeepCopy(),
	}
}

type SubteamLogPoint struct {
	Name  TeamName `codec:"name" json:"name"`
	Seqno Seqno    `codec:"seqno" json:"seqno"`
}

func (o SubteamLogPoint) DeepCopy() SubteamLogPoint {
	return SubteamLogPoint{
		Name:  o.Name.DeepCopy(),
		Seqno: o.Seqno.DeepCopy(),
	}
}

type TeamNamePart string

func (o TeamNamePart) DeepCopy() TeamNamePart {
	return o
}

type TeamName struct {
	Parts []TeamNamePart `codec:"parts" json:"parts"`
}

func (o TeamName) DeepCopy() TeamName {
	return TeamName{
		Parts: (func(x []TeamNamePart) []TeamNamePart {
			if x == nil {
				return nil
			}
			ret := make([]TeamNamePart, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Parts),
	}
}

type TeamCLKRResetUser struct {
	Uid               UID   `codec:"uid" json:"uid"`
	UserEldestSeqno   Seqno `codec:"userEldestSeqno" json:"user_eldest"`
	MemberEldestSeqno Seqno `codec:"memberEldestSeqno" json:"member_eldest"`
}

func (o TeamCLKRResetUser) DeepCopy() TeamCLKRResetUser {
	return TeamCLKRResetUser{
		Uid:               o.Uid.DeepCopy(),
		UserEldestSeqno:   o.UserEldestSeqno.DeepCopy(),
		MemberEldestSeqno: o.MemberEldestSeqno.DeepCopy(),
	}
}

type TeamCLKRMsg struct {
	TeamID              TeamID               `codec:"teamID" json:"team_id"`
	Generation          PerTeamKeyGeneration `codec:"generation" json:"generation"`
	Score               int                  `codec:"score" json:"score"`
	ResetUsersUntrusted []TeamCLKRResetUser  `codec:"resetUsersUntrusted" json:"reset_users"`
}

func (o TeamCLKRMsg) DeepCopy() TeamCLKRMsg {
	return TeamCLKRMsg{
		TeamID:     o.TeamID.DeepCopy(),
		Generation: o.Generation.DeepCopy(),
		Score:      o.Score,
		ResetUsersUntrusted: (func(x []TeamCLKRResetUser) []TeamCLKRResetUser {
			if x == nil {
				return nil
			}
			ret := make([]TeamCLKRResetUser, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ResetUsersUntrusted),
	}
}

type TeamResetUser struct {
	Username    string `codec:"username" json:"username"`
	Uid         UID    `codec:"uid" json:"uid"`
	EldestSeqno Seqno  `codec:"eldestSeqno" json:"eldest_seqno"`
	IsDelete    bool   `codec:"isDelete" json:"is_delete"`
}

func (o TeamResetUser) DeepCopy() TeamResetUser {
	return TeamResetUser{
		Username:    o.Username,
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		IsDelete:    o.IsDelete,
	}
}

type TeamMemberOutFromReset struct {
	TeamID    TeamID        `codec:"teamID" json:"team_id"`
	TeamName  string        `codec:"teamName" json:"team_name"`
	ResetUser TeamResetUser `codec:"resetUser" json:"reset_user"`
}

func (o TeamMemberOutFromReset) DeepCopy() TeamMemberOutFromReset {
	return TeamMemberOutFromReset{
		TeamID:    o.TeamID.DeepCopy(),
		TeamName:  o.TeamName,
		ResetUser: o.ResetUser.DeepCopy(),
	}
}

type TeamChangeRow struct {
	Id                  TeamID `codec:"id" json:"id"`
	Name                string `codec:"name" json:"name"`
	KeyRotated          bool   `codec:"keyRotated" json:"key_rotated"`
	MembershipChanged   bool   `codec:"membershipChanged" json:"membership_changed"`
	LatestSeqno         Seqno  `codec:"latestSeqno" json:"latest_seqno"`
	LatestHiddenSeqno   Seqno  `codec:"latestHiddenSeqno" json:"latest_hidden_seqno"`
	LatestOffchainSeqno Seqno  `codec:"latestOffchainSeqno" json:"latest_offchain_seqno"`
	ImplicitTeam        bool   `codec:"implicitTeam" json:"implicit_team"`
	Misc                bool   `codec:"misc" json:"misc"`
	RemovedResetUsers   bool   `codec:"removedResetUsers" json:"removed_reset_users"`
}

func (o TeamChangeRow) DeepCopy() TeamChangeRow {
	return TeamChangeRow{
		Id:                  o.Id.DeepCopy(),
		Name:                o.Name,
		KeyRotated:          o.KeyRotated,
		MembershipChanged:   o.MembershipChanged,
		LatestSeqno:         o.LatestSeqno.DeepCopy(),
		LatestHiddenSeqno:   o.LatestHiddenSeqno.DeepCopy(),
		LatestOffchainSeqno: o.LatestOffchainSeqno.DeepCopy(),
		ImplicitTeam:        o.ImplicitTeam,
		Misc:                o.Misc,
		RemovedResetUsers:   o.RemovedResetUsers,
	}
}

type TeamExitRow struct {
	Id TeamID `codec:"id" json:"id"`
}

func (o TeamExitRow) DeepCopy() TeamExitRow {
	return TeamExitRow{
		Id: o.Id.DeepCopy(),
	}
}

type TeamNewlyAddedRow struct {
	Id   TeamID `codec:"id" json:"id"`
	Name string `codec:"name" json:"name"`
}

func (o TeamNewlyAddedRow) DeepCopy() TeamNewlyAddedRow {
	return TeamNewlyAddedRow{
		Id:   o.Id.DeepCopy(),
		Name: o.Name,
	}
}

type TeamInvitee struct {
	InviteID    TeamInviteID `codec:"inviteID" json:"invite_id"`
	Uid         UID          `codec:"uid" json:"uid"`
	EldestSeqno Seqno        `codec:"eldestSeqno" json:"eldest_seqno"`
	Role        TeamRole     `codec:"role" json:"role"`
}

func (o TeamInvitee) DeepCopy() TeamInvitee {
	return TeamInvitee{
		InviteID:    o.InviteID.DeepCopy(),
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		Role:        o.Role.DeepCopy(),
	}
}

type TeamSBSMsg struct {
	TeamID   TeamID        `codec:"teamID" json:"team_id"`
	Score    int           `codec:"score" json:"score"`
	Invitees []TeamInvitee `codec:"invitees" json:"invitees"`
}

func (o TeamSBSMsg) DeepCopy() TeamSBSMsg {
	return TeamSBSMsg{
		TeamID: o.TeamID.DeepCopy(),
		Score:  o.Score,
		Invitees: (func(x []TeamInvitee) []TeamInvitee {
			if x == nil {
				return nil
			}
			ret := make([]TeamInvitee, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Invitees),
	}
}

type TeamAccessRequest struct {
	Uid         UID   `codec:"uid" json:"uid"`
	EldestSeqno Seqno `codec:"eldestSeqno" json:"eldest_seqno"`
}

func (o TeamAccessRequest) DeepCopy() TeamAccessRequest {
	return TeamAccessRequest{
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
	}
}

type TeamOpenReqMsg struct {
	TeamID TeamID              `codec:"teamID" json:"team_id"`
	Tars   []TeamAccessRequest `codec:"tars" json:"tars"`
}

func (o TeamOpenReqMsg) DeepCopy() TeamOpenReqMsg {
	return TeamOpenReqMsg{
		TeamID: o.TeamID.DeepCopy(),
		Tars: (func(x []TeamAccessRequest) []TeamAccessRequest {
			if x == nil {
				return nil
			}
			ret := make([]TeamAccessRequest, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Tars),
	}
}

type SeitanAKey string

func (o SeitanAKey) DeepCopy() SeitanAKey {
	return o
}

type SeitanIKey string

func (o SeitanIKey) DeepCopy() SeitanIKey {
	return o
}

type SeitanPubKey string

func (o SeitanPubKey) DeepCopy() SeitanPubKey {
	return o
}

type SeitanIKeyV2 string

func (o SeitanIKeyV2) DeepCopy() SeitanIKeyV2 {
	return o
}

type SeitanKeyAndLabelVersion int

const (
	SeitanKeyAndLabelVersion_V1 SeitanKeyAndLabelVersion = 1
	SeitanKeyAndLabelVersion_V2 SeitanKeyAndLabelVersion = 2
)

func (o SeitanKeyAndLabelVersion) DeepCopy() SeitanKeyAndLabelVersion { return o }

var SeitanKeyAndLabelVersionMap = map[string]SeitanKeyAndLabelVersion{
	"V1": 1,
	"V2": 2,
}

var SeitanKeyAndLabelVersionRevMap = map[SeitanKeyAndLabelVersion]string{
	1: "V1",
	2: "V2",
}

func (e SeitanKeyAndLabelVersion) String() string {
	if v, ok := SeitanKeyAndLabelVersionRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type SeitanKeyAndLabel struct {
	V__  SeitanKeyAndLabelVersion   `codec:"v" json:"v"`
	V1__ *SeitanKeyAndLabelVersion1 `codec:"v1,omitempty" json:"v1,omitempty"`
	V2__ *SeitanKeyAndLabelVersion2 `codec:"v2,omitempty" json:"v2,omitempty"`
}

func (o *SeitanKeyAndLabel) V() (ret SeitanKeyAndLabelVersion, err error) {
	switch o.V__ {
	case SeitanKeyAndLabelVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	case SeitanKeyAndLabelVersion_V2:
		if o.V2__ == nil {
			err = errors.New("unexpected nil value for V2__")
			return ret, err
		}
	}
	return o.V__, nil
}

func (o SeitanKeyAndLabel) V1() (res SeitanKeyAndLabelVersion1) {
	if o.V__ != SeitanKeyAndLabelVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return
	}
	return *o.V1__
}

func (o SeitanKeyAndLabel) V2() (res SeitanKeyAndLabelVersion2) {
	if o.V__ != SeitanKeyAndLabelVersion_V2 {
		panic("wrong case accessed")
	}
	if o.V2__ == nil {
		return
	}
	return *o.V2__
}

func NewSeitanKeyAndLabelWithV1(v SeitanKeyAndLabelVersion1) SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__:  SeitanKeyAndLabelVersion_V1,
		V1__: &v,
	}
}

func NewSeitanKeyAndLabelWithV2(v SeitanKeyAndLabelVersion2) SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__:  SeitanKeyAndLabelVersion_V2,
		V2__: &v,
	}
}

func NewSeitanKeyAndLabelDefault(v SeitanKeyAndLabelVersion) SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__: v,
	}
}

func (o SeitanKeyAndLabel) DeepCopy() SeitanKeyAndLabel {
	return SeitanKeyAndLabel{
		V__: o.V__.DeepCopy(),
		V1__: (func(x *SeitanKeyAndLabelVersion1) *SeitanKeyAndLabelVersion1 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.V1__),
		V2__: (func(x *SeitanKeyAndLabelVersion2) *SeitanKeyAndLabelVersion2 {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.V2__),
	}
}

type SeitanKeyAndLabelVersion1 struct {
	I SeitanIKey     `codec:"i" json:"i"`
	L SeitanKeyLabel `codec:"l" json:"l"`
}

func (o SeitanKeyAndLabelVersion1) DeepCopy() SeitanKeyAndLabelVersion1 {
	return SeitanKeyAndLabelVersion1{
		I: o.I.DeepCopy(),
		L: o.L.DeepCopy(),
	}
}

type SeitanKeyAndLabelVersion2 struct {
	K SeitanPubKey   `codec:"k" json:"k"`
	L SeitanKeyLabel `codec:"l" json:"l"`
}

func (o SeitanKeyAndLabelVersion2) DeepCopy() SeitanKeyAndLabelVersion2 {
	return SeitanKeyAndLabelVersion2{
		K: o.K.DeepCopy(),
		L: o.L.DeepCopy(),
	}
}

type SeitanKeyLabelType int

const (
	SeitanKeyLabelType_SMS SeitanKeyLabelType = 1
)

func (o SeitanKeyLabelType) DeepCopy() SeitanKeyLabelType { return o }

var SeitanKeyLabelTypeMap = map[string]SeitanKeyLabelType{
	"SMS": 1,
}

var SeitanKeyLabelTypeRevMap = map[SeitanKeyLabelType]string{
	1: "SMS",
}

func (e SeitanKeyLabelType) String() string {
	if v, ok := SeitanKeyLabelTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type SeitanKeyLabel struct {
	T__   SeitanKeyLabelType `codec:"t" json:"t"`
	Sms__ *SeitanKeyLabelSms `codec:"sms,omitempty" json:"sms,omitempty"`
}

func (o *SeitanKeyLabel) T() (ret SeitanKeyLabelType, err error) {
	switch o.T__ {
	case SeitanKeyLabelType_SMS:
		if o.Sms__ == nil {
			err = errors.New("unexpected nil value for Sms__")
			return ret, err
		}
	}
	return o.T__, nil
}

func (o SeitanKeyLabel) Sms() (res SeitanKeyLabelSms) {
	if o.T__ != SeitanKeyLabelType_SMS {
		panic("wrong case accessed")
	}
	if o.Sms__ == nil {
		return
	}
	return *o.Sms__
}

func NewSeitanKeyLabelWithSms(v SeitanKeyLabelSms) SeitanKeyLabel {
	return SeitanKeyLabel{
		T__:   SeitanKeyLabelType_SMS,
		Sms__: &v,
	}
}

func NewSeitanKeyLabelDefault(t SeitanKeyLabelType) SeitanKeyLabel {
	return SeitanKeyLabel{
		T__: t,
	}
}

func (o SeitanKeyLabel) DeepCopy() SeitanKeyLabel {
	return SeitanKeyLabel{
		T__: o.T__.DeepCopy(),
		Sms__: (func(x *SeitanKeyLabelSms) *SeitanKeyLabelSms {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Sms__),
	}
}

type SeitanKeyLabelSms struct {
	F string `codec:"f" json:"f"`
	N string `codec:"n" json:"n"`
}

func (o SeitanKeyLabelSms) DeepCopy() SeitanKeyLabelSms {
	return SeitanKeyLabelSms{
		F: o.F,
		N: o.N,
	}
}

type TeamSeitanRequest struct {
	InviteID    TeamInviteID `codec:"inviteID" json:"invite_id"`
	Uid         UID          `codec:"uid" json:"uid"`
	EldestSeqno Seqno        `codec:"eldestSeqno" json:"eldest_seqno"`
	Akey        SeitanAKey   `codec:"akey" json:"akey"`
	Role        TeamRole     `codec:"role" json:"role"`
	UnixCTime   int64        `codec:"unixCTime" json:"ctime"`
}

func (o TeamSeitanRequest) DeepCopy() TeamSeitanRequest {
	return TeamSeitanRequest{
		InviteID:    o.InviteID.DeepCopy(),
		Uid:         o.Uid.DeepCopy(),
		EldestSeqno: o.EldestSeqno.DeepCopy(),
		Akey:        o.Akey.DeepCopy(),
		Role:        o.Role.DeepCopy(),
		UnixCTime:   o.UnixCTime,
	}
}

type TeamSeitanMsg struct {
	TeamID  TeamID              `codec:"teamID" json:"team_id"`
	Seitans []TeamSeitanRequest `codec:"seitans" json:"seitans"`
}

func (o TeamSeitanMsg) DeepCopy() TeamSeitanMsg {
	return TeamSeitanMsg{
		TeamID: o.TeamID.DeepCopy(),
		Seitans: (func(x []TeamSeitanRequest) []TeamSeitanRequest {
			if x == nil {
				return nil
			}
			ret := make([]TeamSeitanRequest, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Seitans),
	}
}

type TeamOpenSweepMsg struct {
	TeamID              TeamID              `codec:"teamID" json:"team_id"`
	ResetUsersUntrusted []TeamCLKRResetUser `codec:"resetUsersUntrusted" json:"reset_users"`
}

func (o TeamOpenSweepMsg) DeepCopy() TeamOpenSweepMsg {
	return TeamOpenSweepMsg{
		TeamID: o.TeamID.DeepCopy(),
		ResetUsersUntrusted: (func(x []TeamCLKRResetUser) []TeamCLKRResetUser {
			if x == nil {
				return nil
			}
			ret := make([]TeamCLKRResetUser, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ResetUsersUntrusted),
	}
}

type TeamKBFSKeyRefresher struct {
	Generation int             `codec:"generation" json:"generation"`
	AppType    TeamApplication `codec:"appType" json:"appType"`
}

func (o TeamKBFSKeyRefresher) DeepCopy() TeamKBFSKeyRefresher {
	return TeamKBFSKeyRefresher{
		Generation: o.Generation,
		AppType:    o.AppType.DeepCopy(),
	}
}

// * TeamRefreshData are needed or wanted data requirements that, if unmet, will cause
// * a refresh of the cache.
type TeamRefreshers struct {
	NeedKeyGeneration                     PerTeamKeyGeneration                       `codec:"needKeyGeneration" json:"needKeyGeneration"`
	NeedApplicationsAtGenerations         map[PerTeamKeyGeneration][]TeamApplication `codec:"needApplicationsAtGenerations" json:"needApplicationsAtGenerations"`
	NeedApplicationsAtGenerationsWithKBFS map[PerTeamKeyGeneration][]TeamApplication `codec:"needApplicationsAtGenerationsWithKBFS" json:"needApplicationsAtGenerationsWithKBFS"`
	WantMembers                           []UserVersion                              `codec:"wantMembers" json:"wantMembers"`
	WantMembersRole                       TeamRole                                   `codec:"wantMembersRole" json:"wantMembersRole"`
	NeedKBFSKeyGeneration                 TeamKBFSKeyRefresher                       `codec:"needKBFSKeyGeneration" json:"needKBFSKeyGeneration"`
}

func (o TeamRefreshers) DeepCopy() TeamRefreshers {
	return TeamRefreshers{
		NeedKeyGeneration: o.NeedKeyGeneration.DeepCopy(),
		NeedApplicationsAtGenerations: (func(x map[PerTeamKeyGeneration][]TeamApplication) map[PerTeamKeyGeneration][]TeamApplication {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration][]TeamApplication, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []TeamApplication) []TeamApplication {
					if x == nil {
						return nil
					}
					ret := make([]TeamApplication, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.NeedApplicationsAtGenerations),
		NeedApplicationsAtGenerationsWithKBFS: (func(x map[PerTeamKeyGeneration][]TeamApplication) map[PerTeamKeyGeneration][]TeamApplication {
			if x == nil {
				return nil
			}
			ret := make(map[PerTeamKeyGeneration][]TeamApplication, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := (func(x []TeamApplication) []TeamApplication {
					if x == nil {
						return nil
					}
					ret := make([]TeamApplication, len(x))
					for i, v := range x {
						vCopy := v.DeepCopy()
						ret[i] = vCopy
					}
					return ret
				})(v)
				ret[kCopy] = vCopy
			}
			return ret
		})(o.NeedApplicationsAtGenerationsWithKBFS),
		WantMembers: (func(x []UserVersion) []UserVersion {
			if x == nil {
				return nil
			}
			ret := make([]UserVersion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.WantMembers),
		WantMembersRole:       o.WantMembersRole.DeepCopy(),
		NeedKBFSKeyGeneration: o.NeedKBFSKeyGeneration.DeepCopy(),
	}
}

type LoadTeamArg struct {
	ID                        TeamID         `codec:"ID" json:"ID"`
	Name                      string         `codec:"name" json:"name"`
	Public                    bool           `codec:"public" json:"public"`
	NeedAdmin                 bool           `codec:"needAdmin" json:"needAdmin"`
	RefreshUIDMapper          bool           `codec:"refreshUIDMapper" json:"refreshUIDMapper"`
	Refreshers                TeamRefreshers `codec:"refreshers" json:"refreshers"`
	ForceFullReload           bool           `codec:"forceFullReload" json:"forceFullReload"`
	ForceRepoll               bool           `codec:"forceRepoll" json:"forceRepoll"`
	StaleOK                   bool           `codec:"staleOK" json:"staleOK"`
	AllowNameLookupBurstCache bool           `codec:"allowNameLookupBurstCache" json:"allowNameLookupBurstCache"`
	SkipNeedHiddenRotateCheck bool           `codec:"skipNeedHiddenRotateCheck" json:"skipNeedHiddenRotateCheck"`
	AuditMode                 AuditMode      `codec:"auditMode" json:"auditMode"`
}

func (o LoadTeamArg) DeepCopy() LoadTeamArg {
	return LoadTeamArg{
		ID:                        o.ID.DeepCopy(),
		Name:                      o.Name,
		Public:                    o.Public,
		NeedAdmin:                 o.NeedAdmin,
		RefreshUIDMapper:          o.RefreshUIDMapper,
		Refreshers:                o.Refreshers.DeepCopy(),
		ForceFullReload:           o.ForceFullReload,
		ForceRepoll:               o.ForceRepoll,
		StaleOK:                   o.StaleOK,
		AllowNameLookupBurstCache: o.AllowNameLookupBurstCache,
		SkipNeedHiddenRotateCheck: o.SkipNeedHiddenRotateCheck,
		AuditMode:                 o.AuditMode.DeepCopy(),
	}
}

type FastTeamLoadArg struct {
	ID                    TeamID                 `codec:"ID" json:"ID"`
	Public                bool                   `codec:"public" json:"public"`
	AssertTeamName        *TeamName              `codec:"assertTeamName,omitempty" json:"assertTeamName,omitempty"`
	Applications          []TeamApplication      `codec:"applications" json:"applications"`
	KeyGenerationsNeeded  []PerTeamKeyGeneration `codec:"keyGenerationsNeeded" json:"keyGenerationsNeeded"`
	NeedLatestKey         bool                   `codec:"needLatestKey" json:"needLatestKey"`
	ForceRefresh          bool                   `codec:"forceRefresh" json:"forceRefresh"`
	HiddenChainIsOptional bool                   `codec:"hiddenChainIsOptional" json:"hiddenChainIsOptional"`
}

func (o FastTeamLoadArg) DeepCopy() FastTeamLoadArg {
	return FastTeamLoadArg{
		ID:     o.ID.DeepCopy(),
		Public: o.Public,
		AssertTeamName: (func(x *TeamName) *TeamName {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.AssertTeamName),
		Applications: (func(x []TeamApplication) []TeamApplication {
			if x == nil {
				return nil
			}
			ret := make([]TeamApplication, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Applications),
		KeyGenerationsNeeded: (func(x []PerTeamKeyGeneration) []PerTeamKeyGeneration {
			if x == nil {
				return nil
			}
			ret := make([]PerTeamKeyGeneration, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.KeyGenerationsNeeded),
		NeedLatestKey:         o.NeedLatestKey,
		ForceRefresh:          o.ForceRefresh,
		HiddenChainIsOptional: o.HiddenChainIsOptional,
	}
}

type FastTeamLoadRes struct {
	Name            TeamName             `codec:"name" json:"name"`
	ApplicationKeys []TeamApplicationKey `codec:"applicationKeys" json:"applicationKeys"`
}

func (o FastTeamLoadRes) DeepCopy() FastTeamLoadRes {
	return FastTeamLoadRes{
		Name: o.Name.DeepCopy(),
		ApplicationKeys: (func(x []TeamApplicationKey) []TeamApplicationKey {
			if x == nil {
				return nil
			}
			ret := make([]TeamApplicationKey, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ApplicationKeys),
	}
}

type ImplicitRole struct {
	Role     TeamRole `codec:"role" json:"role"`
	Ancestor TeamID   `codec:"ancestor" json:"ancestor"`
}

func (o ImplicitRole) DeepCopy() ImplicitRole {
	return ImplicitRole{
		Role:     o.Role.DeepCopy(),
		Ancestor: o.Ancestor.DeepCopy(),
	}
}

type MemberInfo struct {
	UserID              UID           `codec:"userID" json:"uid"`
	TeamID              TeamID        `codec:"teamID" json:"team_id"`
	FqName              string        `codec:"fqName" json:"fq_name"`
	IsImplicitTeam      bool          `codec:"isImplicitTeam" json:"is_implicit_team"`
	IsOpenTeam          bool          `codec:"isOpenTeam" json:"is_open_team"`
	Role                TeamRole      `codec:"role" json:"role"`
	Implicit            *ImplicitRole `codec:"implicit,omitempty" json:"implicit,omitempty"`
	MemberCount         int           `codec:"memberCount" json:"member_count"`
	AllowProfilePromote bool          `codec:"allowProfilePromote" json:"allow_profile_promote"`
	IsMemberShowcased   bool          `codec:"isMemberShowcased" json:"is_member_showcased"`
}

func (o MemberInfo) DeepCopy() MemberInfo {
	return MemberInfo{
		UserID:         o.UserID.DeepCopy(),
		TeamID:         o.TeamID.DeepCopy(),
		FqName:         o.FqName,
		IsImplicitTeam: o.IsImplicitTeam,
		IsOpenTeam:     o.IsOpenTeam,
		Role:           o.Role.DeepCopy(),
		Implicit: (func(x *ImplicitRole) *ImplicitRole {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Implicit),
		MemberCount:         o.MemberCount,
		AllowProfilePromote: o.AllowProfilePromote,
		IsMemberShowcased:   o.IsMemberShowcased,
	}
}

type TeamList struct {
	Teams []MemberInfo `codec:"teams" json:"teams"`
}

func (o TeamList) DeepCopy() TeamList {
	return TeamList{
		Teams: (func(x []MemberInfo) []MemberInfo {
			if x == nil {
				return nil
			}
			ret := make([]MemberInfo, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Teams),
	}
}

type AnnotatedMemberInfo struct {
	UserID              UID              `codec:"userID" json:"uid"`
	TeamID              TeamID           `codec:"teamID" json:"team_id"`
	Username            string           `codec:"username" json:"username"`
	FullName            string           `codec:"fullName" json:"full_name"`
	FqName              string           `codec:"fqName" json:"fq_name"`
	IsImplicitTeam      bool             `codec:"isImplicitTeam" json:"is_implicit_team"`
	ImpTeamDisplayName  string           `codec:"impTeamDisplayName" json:"implicit_team_display_name"`
	IsOpenTeam          bool             `codec:"isOpenTeam" json:"is_open_team"`
	Role                TeamRole         `codec:"role" json:"role"`
	Implicit            *ImplicitRole    `codec:"implicit,omitempty" json:"implicit,omitempty"`
	NeedsPUK            bool             `codec:"needsPUK" json:"needsPUK"`
	MemberCount         int              `codec:"memberCount" json:"member_count"`
	EldestSeqno         Seqno            `codec:"eldestSeqno" json:"member_eldest_seqno"`
	AllowProfilePromote bool             `codec:"allowProfilePromote" json:"allow_profile_promote"`
	IsMemberShowcased   bool             `codec:"isMemberShowcased" json:"is_member_showcased"`
	Status              TeamMemberStatus `codec:"status" json:"status"`
}

func (o AnnotatedMemberInfo) DeepCopy() AnnotatedMemberInfo {
	return AnnotatedMemberInfo{
		UserID:             o.UserID.DeepCopy(),
		TeamID:             o.TeamID.DeepCopy(),
		Username:           o.Username,
		FullName:           o.FullName,
		FqName:             o.FqName,
		IsImplicitTeam:     o.IsImplicitTeam,
		ImpTeamDisplayName: o.ImpTeamDisplayName,
		IsOpenTeam:         o.IsOpenTeam,
		Role:               o.Role.DeepCopy(),
		Implicit: (func(x *ImplicitRole) *ImplicitRole {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Implicit),
		NeedsPUK:            o.NeedsPUK,
		MemberCount:         o.MemberCount,
		EldestSeqno:         o.EldestSeqno.DeepCopy(),
		AllowProfilePromote: o.AllowProfilePromote,
		IsMemberShowcased:   o.IsMemberShowcased,
		Status:              o.Status.DeepCopy(),
	}
}

type AnnotatedTeamList struct {
	Teams                  []AnnotatedMemberInfo                `codec:"teams" json:"teams"`
	AnnotatedActiveInvites map[TeamInviteID]AnnotatedTeamInvite `codec:"annotatedActiveInvites" json:"annotatedActiveInvites"`
}

func (o AnnotatedTeamList) DeepCopy() AnnotatedTeamList {
	return AnnotatedTeamList{
		Teams: (func(x []AnnotatedMemberInfo) []AnnotatedMemberInfo {
			if x == nil {
				return nil
			}
			ret := make([]AnnotatedMemberInfo, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Teams),
		AnnotatedActiveInvites: (func(x map[TeamInviteID]AnnotatedTeamInvite) map[TeamInviteID]AnnotatedTeamInvite {
			if x == nil {
				return nil
			}
			ret := make(map[TeamInviteID]AnnotatedTeamInvite, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.AnnotatedActiveInvites),
	}
}

type TeamAddMemberResult struct {
	Invited     bool  `codec:"invited" json:"invited"`
	User        *User `codec:"user,omitempty" json:"user,omitempty"`
	ChatSending bool  `codec:"chatSending" json:"chatSending"`
}

func (o TeamAddMemberResult) DeepCopy() TeamAddMemberResult {
	return TeamAddMemberResult{
		Invited: o.Invited,
		User: (func(x *User) *User {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.User),
		ChatSending: o.ChatSending,
	}
}

type TeamAddMembersResult struct {
	NotAdded []User `codec:"notAdded" json:"notAdded"`
}

func (o TeamAddMembersResult) DeepCopy() TeamAddMembersResult {
	return TeamAddMembersResult{
		NotAdded: (func(x []User) []User {
			if x == nil {
				return nil
			}
			ret := make([]User, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.NotAdded),
	}
}

type TeamJoinRequest struct {
	Name     string `codec:"name" json:"name"`
	Username string `codec:"username" json:"username"`
}

func (o TeamJoinRequest) DeepCopy() TeamJoinRequest {
	return TeamJoinRequest{
		Name:     o.Name,
		Username: o.Username,
	}
}

type TeamTreeResult struct {
	Entries []TeamTreeEntry `codec:"entries" json:"entries"`
}

func (o TeamTreeResult) DeepCopy() TeamTreeResult {
	return TeamTreeResult{
		Entries: (func(x []TeamTreeEntry) []TeamTreeEntry {
			if x == nil {
				return nil
			}
			ret := make([]TeamTreeEntry, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Entries),
	}
}

type TeamTreeEntry struct {
	Name  TeamName `codec:"name" json:"name"`
	Admin bool     `codec:"admin" json:"admin"`
}

func (o TeamTreeEntry) DeepCopy() TeamTreeEntry {
	return TeamTreeEntry{
		Name:  o.Name.DeepCopy(),
		Admin: o.Admin,
	}
}

type SubteamListEntry struct {
	Name        TeamName `codec:"name" json:"name"`
	TeamID      TeamID   `codec:"teamID" json:"teamID"`
	MemberCount int      `codec:"memberCount" json:"memberCount"`
}

func (o SubteamListEntry) DeepCopy() SubteamListEntry {
	return SubteamListEntry{
		Name:        o.Name.DeepCopy(),
		TeamID:      o.TeamID.DeepCopy(),
		MemberCount: o.MemberCount,
	}
}

type SubteamListResult struct {
	Entries []SubteamListEntry `codec:"entries" json:"entries"`
}

func (o SubteamListResult) DeepCopy() SubteamListResult {
	return SubteamListResult{
		Entries: (func(x []SubteamListEntry) []SubteamListEntry {
			if x == nil {
				return nil
			}
			ret := make([]SubteamListEntry, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Entries),
	}
}

type TeamCreateResult struct {
	TeamID       TeamID `codec:"teamID" json:"teamID"`
	ChatSent     bool   `codec:"chatSent" json:"chatSent"`
	CreatorAdded bool   `codec:"creatorAdded" json:"creatorAdded"`
}

func (o TeamCreateResult) DeepCopy() TeamCreateResult {
	return TeamCreateResult{
		TeamID:       o.TeamID.DeepCopy(),
		ChatSent:     o.ChatSent,
		CreatorAdded: o.CreatorAdded,
	}
}

type TeamSettings struct {
	Open   bool     `codec:"open" json:"open"`
	JoinAs TeamRole `codec:"joinAs" json:"joinAs"`
}

func (o TeamSettings) DeepCopy() TeamSettings {
	return TeamSettings{
		Open:   o.Open,
		JoinAs: o.JoinAs.DeepCopy(),
	}
}

type TeamBotSettings struct {
	Cmds     bool     `codec:"cmds" json:"cmds"`
	Mentions bool     `codec:"mentions" json:"mentions"`
	Triggers []string `codec:"triggers" json:"triggers"`
	Convs    []string `codec:"convs" json:"convs"`
}

func (o TeamBotSettings) DeepCopy() TeamBotSettings {
	return TeamBotSettings{
		Cmds:     o.Cmds,
		Mentions: o.Mentions,
		Triggers: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Triggers),
		Convs: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Convs),
	}
}

type TeamRequestAccessResult struct {
	Open bool `codec:"open" json:"open"`
}

func (o TeamRequestAccessResult) DeepCopy() TeamRequestAccessResult {
	return TeamRequestAccessResult{
		Open: o.Open,
	}
}

type TeamAcceptOrRequestResult struct {
	WasToken    bool `codec:"wasToken" json:"wasToken"`
	WasSeitan   bool `codec:"wasSeitan" json:"wasSeitan"`
	WasTeamName bool `codec:"wasTeamName" json:"wasTeamName"`
	WasOpenTeam bool `codec:"wasOpenTeam" json:"wasOpenTeam"`
}

func (o TeamAcceptOrRequestResult) DeepCopy() TeamAcceptOrRequestResult {
	return TeamAcceptOrRequestResult{
		WasToken:    o.WasToken,
		WasSeitan:   o.WasSeitan,
		WasTeamName: o.WasTeamName,
		WasOpenTeam: o.WasOpenTeam,
	}
}

type TeamShowcase struct {
	IsShowcased       bool    `codec:"isShowcased" json:"is_showcased"`
	Description       *string `codec:"description,omitempty" json:"description,omitempty"`
	SetByUID          *UID    `codec:"setByUID,omitempty" json:"set_by_uid,omitempty"`
	AnyMemberShowcase bool    `codec:"anyMemberShowcase" json:"any_member_showcase"`
}

func (o TeamShowcase) DeepCopy() TeamShowcase {
	return TeamShowcase{
		IsShowcased: o.IsShowcased,
		Description: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Description),
		SetByUID: (func(x *UID) *UID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.SetByUID),
		AnyMemberShowcase: o.AnyMemberShowcase,
	}
}

type TeamAndMemberShowcase struct {
	TeamShowcase      TeamShowcase `codec:"teamShowcase" json:"teamShowcase"`
	IsMemberShowcased bool         `codec:"isMemberShowcased" json:"isMemberShowcased"`
}

func (o TeamAndMemberShowcase) DeepCopy() TeamAndMemberShowcase {
	return TeamAndMemberShowcase{
		TeamShowcase:      o.TeamShowcase.DeepCopy(),
		IsMemberShowcased: o.IsMemberShowcased,
	}
}

type UserRolePair struct {
	AssertionOrEmail string           `codec:"assertionOrEmail" json:"assertionOrEmail"`
	Role             TeamRole         `codec:"role" json:"role"`
	BotSettings      *TeamBotSettings `codec:"botSettings,omitempty" json:"botSettings,omitempty"`
}

func (o UserRolePair) DeepCopy() UserRolePair {
	return UserRolePair{
		AssertionOrEmail: o.AssertionOrEmail,
		Role:             o.Role.DeepCopy(),
		BotSettings: (func(x *TeamBotSettings) *TeamBotSettings {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.BotSettings),
	}
}

type BulkRes struct {
	Invited        []string `codec:"invited" json:"invited"`
	AlreadyInvited []string `codec:"alreadyInvited" json:"alreadyInvited"`
	Malformed      []string `codec:"malformed" json:"malformed"`
}

func (o BulkRes) DeepCopy() BulkRes {
	return BulkRes{
		Invited: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Invited),
		AlreadyInvited: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.AlreadyInvited),
		Malformed: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Malformed),
	}
}

type ImplicitTeamUserSet struct {
	KeybaseUsers    []string          `codec:"keybaseUsers" json:"keybaseUsers"`
	UnresolvedUsers []SocialAssertion `codec:"unresolvedUsers" json:"unresolvedUsers"`
}

func (o ImplicitTeamUserSet) DeepCopy() ImplicitTeamUserSet {
	return ImplicitTeamUserSet{
		KeybaseUsers: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.KeybaseUsers),
		UnresolvedUsers: (func(x []SocialAssertion) []SocialAssertion {
			if x == nil {
				return nil
			}
			ret := make([]SocialAssertion, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.UnresolvedUsers),
	}
}

// * iTeams
type ImplicitTeamDisplayName struct {
	IsPublic     bool                      `codec:"isPublic" json:"isPublic"`
	Writers      ImplicitTeamUserSet       `codec:"writers" json:"writers"`
	Readers      ImplicitTeamUserSet       `codec:"readers" json:"readers"`
	ConflictInfo *ImplicitTeamConflictInfo `codec:"conflictInfo,omitempty" json:"conflictInfo,omitempty"`
}

func (o ImplicitTeamDisplayName) DeepCopy() ImplicitTeamDisplayName {
	return ImplicitTeamDisplayName{
		IsPublic: o.IsPublic,
		Writers:  o.Writers.DeepCopy(),
		Readers:  o.Readers.DeepCopy(),
		ConflictInfo: (func(x *ImplicitTeamConflictInfo) *ImplicitTeamConflictInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ConflictInfo),
	}
}

type ConflictGeneration int

func (o ConflictGeneration) DeepCopy() ConflictGeneration {
	return o
}

type ImplicitTeamConflictInfo struct {
	Generation ConflictGeneration `codec:"generation" json:"generation"`
	Time       Time               `codec:"time" json:"time"`
}

func (o ImplicitTeamConflictInfo) DeepCopy() ImplicitTeamConflictInfo {
	return ImplicitTeamConflictInfo{
		Generation: o.Generation.DeepCopy(),
		Time:       o.Time.DeepCopy(),
	}
}

type LookupImplicitTeamRes struct {
	TeamID      TeamID                  `codec:"teamID" json:"teamID"`
	Name        TeamName                `codec:"name" json:"name"`
	DisplayName ImplicitTeamDisplayName `codec:"displayName" json:"displayName"`
	TlfID       TLFID                   `codec:"tlfID" json:"tlfID"`
}

func (o LookupImplicitTeamRes) DeepCopy() LookupImplicitTeamRes {
	return LookupImplicitTeamRes{
		TeamID:      o.TeamID.DeepCopy(),
		Name:        o.Name.DeepCopy(),
		DisplayName: o.DisplayName.DeepCopy(),
		TlfID:       o.TlfID.DeepCopy(),
	}
}

type TeamOperation struct {
	ManageMembers          bool `codec:"manageMembers" json:"manageMembers"`
	ManageSubteams         bool `codec:"manageSubteams" json:"manageSubteams"`
	CreateChannel          bool `codec:"createChannel" json:"createChannel"`
	Chat                   bool `codec:"chat" json:"chat"`
	DeleteChannel          bool `codec:"deleteChannel" json:"deleteChannel"`
	RenameChannel          bool `codec:"renameChannel" json:"renameChannel"`
	RenameTeam             bool `codec:"renameTeam" json:"renameTeam"`
	EditChannelDescription bool `codec:"editChannelDescription" json:"editChannelDescription"`
	EditTeamDescription    bool `codec:"editTeamDescription" json:"editTeamDescription"`
	SetTeamShowcase        bool `codec:"setTeamShowcase" json:"setTeamShowcase"`
	SetMemberShowcase      bool `codec:"setMemberShowcase" json:"setMemberShowcase"`
	SetRetentionPolicy     bool `codec:"setRetentionPolicy" json:"setRetentionPolicy"`
	SetMinWriterRole       bool `codec:"setMinWriterRole" json:"setMinWriterRole"`
	ChangeOpenTeam         bool `codec:"changeOpenTeam" json:"changeOpenTeam"`
	LeaveTeam              bool `codec:"leaveTeam" json:"leaveTeam"`
	JoinTeam               bool `codec:"joinTeam" json:"joinTeam"`
	SetPublicityAny        bool `codec:"setPublicityAny" json:"setPublicityAny"`
	ListFirst              bool `codec:"listFirst" json:"listFirst"`
	ChangeTarsDisabled     bool `codec:"changeTarsDisabled" json:"changeTarsDisabled"`
	DeleteChatHistory      bool `codec:"deleteChatHistory" json:"deleteChatHistory"`
	DeleteOtherMessages    bool `codec:"deleteOtherMessages" json:"deleteOtherMessages"`
	DeleteTeam             bool `codec:"deleteTeam" json:"deleteTeam"`
	PinMessage             bool `codec:"pinMessage" json:"pinMessage"`
	ManageBots             bool `codec:"manageBots" json:"manageBots"`
}

func (o TeamOperation) DeepCopy() TeamOperation {
	return TeamOperation{
		ManageMembers:          o.ManageMembers,
		ManageSubteams:         o.ManageSubteams,
		CreateChannel:          o.CreateChannel,
		Chat:                   o.Chat,
		DeleteChannel:          o.DeleteChannel,
		RenameChannel:          o.RenameChannel,
		RenameTeam:             o.RenameTeam,
		EditChannelDescription: o.EditChannelDescription,
		EditTeamDescription:    o.EditTeamDescription,
		SetTeamShowcase:        o.SetTeamShowcase,
		SetMemberShowcase:      o.SetMemberShowcase,
		SetRetentionPolicy:     o.SetRetentionPolicy,
		SetMinWriterRole:       o.SetMinWriterRole,
		ChangeOpenTeam:         o.ChangeOpenTeam,
		LeaveTeam:              o.LeaveTeam,
		JoinTeam:               o.JoinTeam,
		SetPublicityAny:        o.SetPublicityAny,
		ListFirst:              o.ListFirst,
		ChangeTarsDisabled:     o.ChangeTarsDisabled,
		DeleteChatHistory:      o.DeleteChatHistory,
		DeleteOtherMessages:    o.DeleteOtherMessages,
		DeleteTeam:             o.DeleteTeam,
		PinMessage:             o.PinMessage,
		ManageBots:             o.ManageBots,
	}
}

type ProfileTeamLoadRes struct {
	LoadTimeNsec int64 `codec:"loadTimeNsec" json:"loadTimeNsec"`
}

func (o ProfileTeamLoadRes) DeepCopy() ProfileTeamLoadRes {
	return ProfileTeamLoadRes{
		LoadTimeNsec: o.LoadTimeNsec,
	}
}

type RotationType int

const (
	RotationType_VISIBLE RotationType = 0
	RotationType_HIDDEN  RotationType = 1
	RotationType_CLKR    RotationType = 2
)

func (o RotationType) DeepCopy() RotationType { return o }

var RotationTypeMap = map[string]RotationType{
	"VISIBLE": 0,
	"HIDDEN":  1,
	"CLKR":    2,
}

var RotationTypeRevMap = map[RotationType]string{
	0: "VISIBLE",
	1: "HIDDEN",
	2: "CLKR",
}

func (e RotationType) String() string {
	if v, ok := RotationTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type TeamDebugRes struct {
	Chain TeamSigChainState `codec:"chain" json:"chain"`
}

func (o TeamDebugRes) DeepCopy() TeamDebugRes {
	return TeamDebugRes{
		Chain: o.Chain.DeepCopy(),
	}
}

type TeamProfileAddEntry struct {
	TeamID         TeamID   `codec:"teamID" json:"teamID"`
	TeamName       TeamName `codec:"teamName" json:"teamName"`
	Open           bool     `codec:"open" json:"open"`
	DisabledReason string   `codec:"disabledReason" json:"disabledReason"`
}

func (o TeamProfileAddEntry) DeepCopy() TeamProfileAddEntry {
	return TeamProfileAddEntry{
		TeamID:         o.TeamID.DeepCopy(),
		TeamName:       o.TeamName.DeepCopy(),
		Open:           o.Open,
		DisabledReason: o.DisabledReason,
	}
}

type MemberEmail struct {
	Email string `codec:"email" json:"email"`
	Role  string `codec:"role" json:"role"`
}

func (o MemberEmail) DeepCopy() MemberEmail {
	return MemberEmail{
		Email: o.Email,
		Role:  o.Role,
	}
}

type MemberUsername struct {
	Username string `codec:"username" json:"username"`
	Role     string `codec:"role" json:"role"`
}

func (o MemberUsername) DeepCopy() MemberUsername {
	return MemberUsername{
		Username: o.Username,
		Role:     o.Role,
	}
}

type TeamRolePair struct {
	Role         TeamRole `codec:"role" json:"role"`
	ImplicitRole TeamRole `codec:"implicitRole" json:"implicit_role"`
}

func (o TeamRolePair) DeepCopy() TeamRolePair {
	return TeamRolePair{
		Role:         o.Role.DeepCopy(),
		ImplicitRole: o.ImplicitRole.DeepCopy(),
	}
}

type TeamRoleMapAndVersion struct {
	Teams   map[TeamID]TeamRolePair `codec:"teams" json:"teams"`
	Version UserTeamVersion         `codec:"version" json:"user_team_version"`
}

func (o TeamRoleMapAndVersion) DeepCopy() TeamRoleMapAndVersion {
	return TeamRoleMapAndVersion{
		Teams: (func(x map[TeamID]TeamRolePair) map[TeamID]TeamRolePair {
			if x == nil {
				return nil
			}
			ret := make(map[TeamID]TeamRolePair, len(x))
			for k, v := range x {
				kCopy := k.DeepCopy()
				vCopy := v.DeepCopy()
				ret[kCopy] = vCopy
			}
			return ret
		})(o.Teams),
		Version: o.Version.DeepCopy(),
	}
}

type TeamRoleMapStored struct {
	Data     TeamRoleMapAndVersion `codec:"data" json:"data"`
	CachedAt Time                  `codec:"cachedAt" json:"cachedAt"`
}

func (o TeamRoleMapStored) DeepCopy() TeamRoleMapStored {
	return TeamRoleMapStored{
		Data:     o.Data.DeepCopy(),
		CachedAt: o.CachedAt.DeepCopy(),
	}
}

type UserTeamVersion int

func (o UserTeamVersion) DeepCopy() UserTeamVersion {
	return o
}

type UserTeamVersionUpdate struct {
	Version UserTeamVersion `codec:"version" json:"version"`
}

func (o UserTeamVersionUpdate) DeepCopy() UserTeamVersionUpdate {
	return UserTeamVersionUpdate{
		Version: o.Version.DeepCopy(),
	}
}

type AnnotatedTeamMemberDetails struct {
	Details TeamMemberDetails `codec:"details" json:"details"`
	Role    TeamRole          `codec:"role" json:"role"`
}

func (o AnnotatedTeamMemberDetails) DeepCopy() AnnotatedTeamMemberDetails {
	return AnnotatedTeamMemberDetails{
		Details: o.Details.DeepCopy(),
		Role:    o.Role.DeepCopy(),
	}
}

type AnnotatedTeam struct {
	TeamID                       TeamID                       `codec:"teamID" json:"teamID"`
	Name                         string                       `codec:"name" json:"name"`
	TransitiveSubteamsUnverified SubteamListResult            `codec:"transitiveSubteamsUnverified" json:"transitiveSubteamsUnverified"`
	Members                      []AnnotatedTeamMemberDetails `codec:"members" json:"members"`
	Invites                      []AnnotatedTeamInvite        `codec:"invites" json:"invites"`
	JoinRequests                 []TeamJoinRequest            `codec:"joinRequests" json:"joinRequests"`
	UserIsShowcasing             bool                         `codec:"userIsShowcasing" json:"userIsShowcasing"`
	TarsDisabled                 bool                         `codec:"tarsDisabled" json:"tarsDisabled"`
	Settings                     TeamSettings                 `codec:"settings" json:"settings"`
	Showcase                     TeamShowcase                 `codec:"showcase" json:"showcase"`
}

func (o AnnotatedTeam) DeepCopy() AnnotatedTeam {
	return AnnotatedTeam{
		TeamID:                       o.TeamID.DeepCopy(),
		Name:                         o.Name,
		TransitiveSubteamsUnverified: o.TransitiveSubteamsUnverified.DeepCopy(),
		Members: (func(x []AnnotatedTeamMemberDetails) []AnnotatedTeamMemberDetails {
			if x == nil {
				return nil
			}
			ret := make([]AnnotatedTeamMemberDetails, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Members),
		Invites: (func(x []AnnotatedTeamInvite) []AnnotatedTeamInvite {
			if x == nil {
				return nil
			}
			ret := make([]AnnotatedTeamInvite, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Invites),
		JoinRequests: (func(x []TeamJoinRequest) []TeamJoinRequest {
			if x == nil {
				return nil
			}
			ret := make([]TeamJoinRequest, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.JoinRequests),
		UserIsShowcasing: o.UserIsShowcasing,
		TarsDisabled:     o.TarsDisabled,
		Settings:         o.Settings.DeepCopy(),
		Showcase:         o.Showcase.DeepCopy(),
	}
}