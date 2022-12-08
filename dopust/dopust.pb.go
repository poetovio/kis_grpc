// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: dopust.proto

package dopust

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Dopust struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdDopust        int64  `protobuf:"varint,1,opt,name=id_dopust,json=idDopust,proto3" json:"id_dopust,omitempty"`
	IdZaposlenega   int64  `protobuf:"varint,2,opt,name=id_zaposlenega,json=idZaposlenega,proto3" json:"id_zaposlenega,omitempty"`
	DatumZacetka    string `protobuf:"bytes,3,opt,name=datumZacetka,proto3" json:"datumZacetka,omitempty"`
	DatumKonca      string `protobuf:"bytes,4,opt,name=datumKonca,proto3" json:"datumKonca,omitempty"`
	VrstaDopusta    string `protobuf:"bytes,5,opt,name=vrstaDopusta,proto3" json:"vrstaDopusta,omitempty"`
	StevilkaDopusta int64  `protobuf:"varint,6,opt,name=stevilkaDopusta,proto3" json:"stevilkaDopusta,omitempty"`
}

func (x *Dopust) Reset() {
	*x = Dopust{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dopust) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dopust) ProtoMessage() {}

func (x *Dopust) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dopust.ProtoReflect.Descriptor instead.
func (*Dopust) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{0}
}

func (x *Dopust) GetIdDopust() int64 {
	if x != nil {
		return x.IdDopust
	}
	return 0
}

func (x *Dopust) GetIdZaposlenega() int64 {
	if x != nil {
		return x.IdZaposlenega
	}
	return 0
}

func (x *Dopust) GetDatumZacetka() string {
	if x != nil {
		return x.DatumZacetka
	}
	return ""
}

func (x *Dopust) GetDatumKonca() string {
	if x != nil {
		return x.DatumKonca
	}
	return ""
}

func (x *Dopust) GetVrstaDopusta() string {
	if x != nil {
		return x.VrstaDopusta
	}
	return ""
}

func (x *Dopust) GetStevilkaDopusta() int64 {
	if x != nil {
		return x.StevilkaDopusta
	}
	return 0
}

type Zaposlen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdZaposlenega int64  `protobuf:"varint,1,opt,name=id_zaposlenega,json=idZaposlenega,proto3" json:"id_zaposlenega,omitempty"`
	Ime           string `protobuf:"bytes,2,opt,name=ime,proto3" json:"ime,omitempty"`
	Priimek       string `protobuf:"bytes,3,opt,name=priimek,proto3" json:"priimek,omitempty"`
	Spol          string `protobuf:"bytes,4,opt,name=spol,proto3" json:"spol,omitempty"`
	DatumRojstva  string `protobuf:"bytes,5,opt,name=datumRojstva,proto3" json:"datumRojstva,omitempty"`
	Podjetje      string `protobuf:"bytes,6,opt,name=podjetje,proto3" json:"podjetje,omitempty"`
}

func (x *Zaposlen) Reset() {
	*x = Zaposlen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zaposlen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zaposlen) ProtoMessage() {}

func (x *Zaposlen) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zaposlen.ProtoReflect.Descriptor instead.
func (*Zaposlen) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{1}
}

func (x *Zaposlen) GetIdZaposlenega() int64 {
	if x != nil {
		return x.IdZaposlenega
	}
	return 0
}

func (x *Zaposlen) GetIme() string {
	if x != nil {
		return x.Ime
	}
	return ""
}

func (x *Zaposlen) GetPriimek() string {
	if x != nil {
		return x.Priimek
	}
	return ""
}

func (x *Zaposlen) GetSpol() string {
	if x != nil {
		return x.Spol
	}
	return ""
}

func (x *Zaposlen) GetDatumRojstva() string {
	if x != nil {
		return x.DatumRojstva
	}
	return ""
}

func (x *Zaposlen) GetPodjetje() string {
	if x != nil {
		return x.Podjetje
	}
	return ""
}

type GetZaposlen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdZaposlen int64 `protobuf:"varint,1,opt,name=id_zaposlen,json=idZaposlen,proto3" json:"id_zaposlen,omitempty"`
}

func (x *GetZaposlen) Reset() {
	*x = GetZaposlen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZaposlen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZaposlen) ProtoMessage() {}

func (x *GetZaposlen) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZaposlen.ProtoReflect.Descriptor instead.
func (*GetZaposlen) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{2}
}

func (x *GetZaposlen) GetIdZaposlen() int64 {
	if x != nil {
		return x.IdZaposlen
	}
	return 0
}

type CreateZaposlenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ime          string `protobuf:"bytes,1,opt,name=ime,proto3" json:"ime,omitempty"`
	Priimek      string `protobuf:"bytes,2,opt,name=priimek,proto3" json:"priimek,omitempty"`
	Spol         string `protobuf:"bytes,3,opt,name=spol,proto3" json:"spol,omitempty"`
	DatumRojstva string `protobuf:"bytes,4,opt,name=datumRojstva,proto3" json:"datumRojstva,omitempty"`
	Podjetje     string `protobuf:"bytes,5,opt,name=podjetje,proto3" json:"podjetje,omitempty"`
}

func (x *CreateZaposlenRequest) Reset() {
	*x = CreateZaposlenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZaposlenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZaposlenRequest) ProtoMessage() {}

func (x *CreateZaposlenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZaposlenRequest.ProtoReflect.Descriptor instead.
func (*CreateZaposlenRequest) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{3}
}

func (x *CreateZaposlenRequest) GetIme() string {
	if x != nil {
		return x.Ime
	}
	return ""
}

func (x *CreateZaposlenRequest) GetPriimek() string {
	if x != nil {
		return x.Priimek
	}
	return ""
}

func (x *CreateZaposlenRequest) GetSpol() string {
	if x != nil {
		return x.Spol
	}
	return ""
}

func (x *CreateZaposlenRequest) GetDatumRojstva() string {
	if x != nil {
		return x.DatumRojstva
	}
	return ""
}

func (x *CreateZaposlenRequest) GetPodjetje() string {
	if x != nil {
		return x.Podjetje
	}
	return ""
}

type GetDopustRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdDopust int64 `protobuf:"varint,1,opt,name=id_dopust,json=idDopust,proto3" json:"id_dopust,omitempty"`
}

func (x *GetDopustRequest) Reset() {
	*x = GetDopustRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDopustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDopustRequest) ProtoMessage() {}

func (x *GetDopustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDopustRequest.ProtoReflect.Descriptor instead.
func (*GetDopustRequest) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{4}
}

func (x *GetDopustRequest) GetIdDopust() int64 {
	if x != nil {
		return x.IdDopust
	}
	return 0
}

type UpdateDopustRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdZaposlenega   int64  `protobuf:"varint,2,opt,name=id_zaposlenega,json=idZaposlenega,proto3" json:"id_zaposlenega,omitempty"`
	DatumZacetka    string `protobuf:"bytes,3,opt,name=datumZacetka,proto3" json:"datumZacetka,omitempty"`
	DatumKonca      string `protobuf:"bytes,4,opt,name=datumKonca,proto3" json:"datumKonca,omitempty"`
	VrstaDopusta    string `protobuf:"bytes,5,opt,name=vrstaDopusta,proto3" json:"vrstaDopusta,omitempty"`
	StevilkaDopusta int64  `protobuf:"varint,6,opt,name=stevilkaDopusta,proto3" json:"stevilkaDopusta,omitempty"`
}

func (x *UpdateDopustRequest) Reset() {
	*x = UpdateDopustRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDopustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDopustRequest) ProtoMessage() {}

func (x *UpdateDopustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDopustRequest.ProtoReflect.Descriptor instead.
func (*UpdateDopustRequest) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDopustRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDopustRequest) GetIdZaposlenega() int64 {
	if x != nil {
		return x.IdZaposlenega
	}
	return 0
}

func (x *UpdateDopustRequest) GetDatumZacetka() string {
	if x != nil {
		return x.DatumZacetka
	}
	return ""
}

func (x *UpdateDopustRequest) GetDatumKonca() string {
	if x != nil {
		return x.DatumKonca
	}
	return ""
}

func (x *UpdateDopustRequest) GetVrstaDopusta() string {
	if x != nil {
		return x.VrstaDopusta
	}
	return ""
}

func (x *UpdateDopustRequest) GetStevilkaDopusta() int64 {
	if x != nil {
		return x.StevilkaDopusta
	}
	return 0
}

type CreateDopustRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdZaposlenega   int64  `protobuf:"varint,1,opt,name=id_zaposlenega,json=idZaposlenega,proto3" json:"id_zaposlenega,omitempty"`
	DatumZacetka    string `protobuf:"bytes,2,opt,name=datumZacetka,proto3" json:"datumZacetka,omitempty"`
	DatumKonca      string `protobuf:"bytes,3,opt,name=datumKonca,proto3" json:"datumKonca,omitempty"`
	VrstaDopusta    string `protobuf:"bytes,4,opt,name=vrstaDopusta,proto3" json:"vrstaDopusta,omitempty"`
	StevilkaDopusta int64  `protobuf:"varint,5,opt,name=stevilkaDopusta,proto3" json:"stevilkaDopusta,omitempty"`
}

func (x *CreateDopustRequest) Reset() {
	*x = CreateDopustRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDopustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDopustRequest) ProtoMessage() {}

func (x *CreateDopustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDopustRequest.ProtoReflect.Descriptor instead.
func (*CreateDopustRequest) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{6}
}

func (x *CreateDopustRequest) GetIdZaposlenega() int64 {
	if x != nil {
		return x.IdZaposlenega
	}
	return 0
}

func (x *CreateDopustRequest) GetDatumZacetka() string {
	if x != nil {
		return x.DatumZacetka
	}
	return ""
}

func (x *CreateDopustRequest) GetDatumKonca() string {
	if x != nil {
		return x.DatumKonca
	}
	return ""
}

func (x *CreateDopustRequest) GetVrstaDopusta() string {
	if x != nil {
		return x.VrstaDopusta
	}
	return ""
}

func (x *CreateDopustRequest) GetStevilkaDopusta() int64 {
	if x != nil {
		return x.StevilkaDopusta
	}
	return 0
}

type DeleteDopustResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteDopustResponse) Reset() {
	*x = DeleteDopustResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDopustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDopustResponse) ProtoMessage() {}

func (x *DeleteDopustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDopustResponse.ProtoReflect.Descriptor instead.
func (*DeleteDopustResponse) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDopustResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetDopustiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *int64 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit *int64 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetDopustiRequest) Reset() {
	*x = GetDopustiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDopustiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDopustiRequest) ProtoMessage() {}

func (x *GetDopustiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDopustiRequest.ProtoReflect.Descriptor instead.
func (*GetDopustiRequest) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{8}
}

func (x *GetDopustiRequest) GetPage() int64 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *GetDopustiRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetZaposleniParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetZaposleniParams) Reset() {
	*x = GetZaposleniParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZaposleniParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZaposleniParams) ProtoMessage() {}

func (x *GetZaposleniParams) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZaposleniParams.ProtoReflect.Descriptor instead.
func (*GetZaposleniParams) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{9}
}

type ZaposlenList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zaposleni []*Zaposlen `protobuf:"bytes,1,rep,name=zaposleni,proto3" json:"zaposleni,omitempty"`
}

func (x *ZaposlenList) Reset() {
	*x = ZaposlenList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZaposlenList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZaposlenList) ProtoMessage() {}

func (x *ZaposlenList) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZaposlenList.ProtoReflect.Descriptor instead.
func (*ZaposlenList) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{10}
}

func (x *ZaposlenList) GetZaposleni() []*Zaposlen {
	if x != nil {
		return x.Zaposleni
	}
	return nil
}

type GetDopustiParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDopustiParams) Reset() {
	*x = GetDopustiParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDopustiParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDopustiParams) ProtoMessage() {}

func (x *GetDopustiParams) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDopustiParams.ProtoReflect.Descriptor instead.
func (*GetDopustiParams) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{11}
}

type DopustList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dopusti []*Dopust `protobuf:"bytes,1,rep,name=dopusti,proto3" json:"dopusti,omitempty"`
}

func (x *DopustList) Reset() {
	*x = DopustList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dopust_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DopustList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DopustList) ProtoMessage() {}

func (x *DopustList) ProtoReflect() protoreflect.Message {
	mi := &file_dopust_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DopustList.ProtoReflect.Descriptor instead.
func (*DopustList) Descriptor() ([]byte, []int) {
	return file_dopust_proto_rawDescGZIP(), []int{12}
}

func (x *DopustList) GetDopusti() []*Dopust {
	if x != nil {
		return x.Dopusti
	}
	return nil
}

var File_dopust_proto protoreflect.FileDescriptor

var file_dopust_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x06, 0x44, 0x6f, 0x70, 0x75, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64, 0x5f, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x7a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x65, 0x67, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x64, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c,
	0x65, 0x6e, 0x65, 0x67, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x5a, 0x61,
	0x63, 0x65, 0x74, 0x6b, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x75, 0x6d, 0x5a, 0x61, 0x63, 0x65, 0x74, 0x6b, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x75, 0x6d, 0x4b, 0x6f, 0x6e, 0x63, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x75, 0x6d, 0x4b, 0x6f, 0x6e, 0x63, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x72, 0x73,
	0x74, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x76, 0x72, 0x73, 0x74, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x12, 0x28, 0x0a,
	0x0f, 0x73, 0x74, 0x65, 0x76, 0x69, 0x6c, 0x6b, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x65, 0x76, 0x69, 0x6c, 0x6b, 0x61,
	0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x22, 0xb1, 0x01, 0x0a, 0x08, 0x5a, 0x61, 0x70, 0x6f,
	0x73, 0x6c, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x7a, 0x61, 0x70, 0x6f, 0x73,
	0x6c, 0x65, 0x6e, 0x65, 0x67, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x64,
	0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x65, 0x67, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x69, 0x69, 0x6d, 0x65, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x69, 0x69, 0x6d, 0x65, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x70, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x64,
	0x61, 0x74, 0x75, 0x6d, 0x52, 0x6f, 0x6a, 0x73, 0x74, 0x76, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x6f, 0x6a, 0x73, 0x74, 0x76, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x64, 0x6a, 0x65, 0x74, 0x6a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6f, 0x64, 0x6a, 0x65, 0x74, 0x6a, 0x65, 0x22, 0x2e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x64,
	0x5f, 0x7a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x69, 0x64, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x69, 0x6d,
	0x65, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x69, 0x6d, 0x65,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x70, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x70, 0x6f, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x6f,
	0x6a, 0x73, 0x74, 0x76, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x75, 0x6d, 0x52, 0x6f, 0x6a, 0x73, 0x74, 0x76, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x64,
	0x6a, 0x65, 0x74, 0x6a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x64,
	0x6a, 0x65, 0x74, 0x6a, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x70, 0x75,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64, 0x5f,
	0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x64,
	0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x7a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x65, 0x67, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x64, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c,
	0x65, 0x6e, 0x65, 0x67, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x5a, 0x61,
	0x63, 0x65, 0x74, 0x6b, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x75, 0x6d, 0x5a, 0x61, 0x63, 0x65, 0x74, 0x6b, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x75, 0x6d, 0x4b, 0x6f, 0x6e, 0x63, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x75, 0x6d, 0x4b, 0x6f, 0x6e, 0x63, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x72, 0x73,
	0x74, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x76, 0x72, 0x73, 0x74, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x12, 0x28, 0x0a,
	0x0f, 0x73, 0x74, 0x65, 0x76, 0x69, 0x6c, 0x6b, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x65, 0x76, 0x69, 0x6c, 0x6b, 0x61,
	0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x22, 0xce, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x64, 0x5f, 0x7a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x65, 0x67,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x64, 0x5a, 0x61, 0x70, 0x6f, 0x73,
	0x6c, 0x65, 0x6e, 0x65, 0x67, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x5a,
	0x61, 0x63, 0x65, 0x74, 0x6b, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61,
	0x74, 0x75, 0x6d, 0x5a, 0x61, 0x63, 0x65, 0x74, 0x6b, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61,
	0x74, 0x75, 0x6d, 0x4b, 0x6f, 0x6e, 0x63, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x61, 0x74, 0x75, 0x6d, 0x4b, 0x6f, 0x6e, 0x63, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x72,
	0x73, 0x74, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x76, 0x72, 0x73, 0x74, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x12, 0x28,
	0x0a, 0x0f, 0x73, 0x74, 0x65, 0x76, 0x69, 0x6c, 0x6b, 0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x65, 0x76, 0x69, 0x6c, 0x6b,
	0x61, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5a, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x5a, 0x61, 0x70,
	0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x69, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x3e, 0x0a, 0x0c,
	0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09,
	0x7a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x69, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65,
	0x6e, 0x52, 0x09, 0x7a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x69, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x69, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x22, 0x36, 0x0a, 0x0a, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x07, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x69, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x52,
	0x07, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x69, 0x32, 0xd6, 0x03, 0x0a, 0x0d, 0x44, 0x6f, 0x70,
	0x75, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x64, 0x6f, 0x70,
	0x75, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74,
	0x2e, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74,
	0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x69,
	0x12, 0x19, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x70,
	0x75, 0x73, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x64, 0x6f,
	0x70, 0x75, 0x73, 0x74, 0x2e, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x3d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x12,
	0x1b, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x70, 0x75, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x64,
	0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x22, 0x00, 0x12, 0x48,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x12, 0x18,
	0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x70, 0x75, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x6f, 0x70, 0x75, 0x73,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x64, 0x6f, 0x70,
	0x75, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x64, 0x6f, 0x70, 0x75,
	0x73, 0x74, 0x2e, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x69, 0x12, 0x1a, 0x2e,
	0x64, 0x6f, 0x70, 0x75, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c,
	0x65, 0x6e, 0x69, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x14, 0x2e, 0x64, 0x6f, 0x70, 0x75,
	0x73, 0x74, 0x2e, 0x5a, 0x61, 0x70, 0x6f, 0x73, 0x6c, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dopust_proto_rawDescOnce sync.Once
	file_dopust_proto_rawDescData = file_dopust_proto_rawDesc
)

func file_dopust_proto_rawDescGZIP() []byte {
	file_dopust_proto_rawDescOnce.Do(func() {
		file_dopust_proto_rawDescData = protoimpl.X.CompressGZIP(file_dopust_proto_rawDescData)
	})
	return file_dopust_proto_rawDescData
}

var file_dopust_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_dopust_proto_goTypes = []interface{}{
	(*Dopust)(nil),                // 0: dopust.Dopust
	(*Zaposlen)(nil),              // 1: dopust.Zaposlen
	(*GetZaposlen)(nil),           // 2: dopust.GetZaposlen
	(*CreateZaposlenRequest)(nil), // 3: dopust.CreateZaposlenRequest
	(*GetDopustRequest)(nil),      // 4: dopust.GetDopustRequest
	(*UpdateDopustRequest)(nil),   // 5: dopust.UpdateDopustRequest
	(*CreateDopustRequest)(nil),   // 6: dopust.CreateDopustRequest
	(*DeleteDopustResponse)(nil),  // 7: dopust.DeleteDopustResponse
	(*GetDopustiRequest)(nil),     // 8: dopust.GetDopustiRequest
	(*GetZaposleniParams)(nil),    // 9: dopust.GetZaposleniParams
	(*ZaposlenList)(nil),          // 10: dopust.ZaposlenList
	(*GetDopustiParams)(nil),      // 11: dopust.GetDopustiParams
	(*DopustList)(nil),            // 12: dopust.DopustList
}
var file_dopust_proto_depIdxs = []int32{
	1,  // 0: dopust.ZaposlenList.zaposleni:type_name -> dopust.Zaposlen
	0,  // 1: dopust.DopustList.dopusti:type_name -> dopust.Dopust
	6,  // 2: dopust.DopustService.CreateDopust:input_type -> dopust.CreateDopustRequest
	4,  // 3: dopust.DopustService.GetDopust:input_type -> dopust.GetDopustRequest
	8,  // 4: dopust.DopustService.GetDopusti:input_type -> dopust.GetDopustiRequest
	5,  // 5: dopust.DopustService.UpdateDopust:input_type -> dopust.UpdateDopustRequest
	4,  // 6: dopust.DopustService.DeleteDopust:input_type -> dopust.GetDopustRequest
	3,  // 7: dopust.DopustService.CreateZaposlen:input_type -> dopust.CreateZaposlenRequest
	9,  // 8: dopust.DopustService.GetZaposleni:input_type -> dopust.GetZaposleniParams
	0,  // 9: dopust.DopustService.CreateDopust:output_type -> dopust.Dopust
	0,  // 10: dopust.DopustService.GetDopust:output_type -> dopust.Dopust
	0,  // 11: dopust.DopustService.GetDopusti:output_type -> dopust.Dopust
	0,  // 12: dopust.DopustService.UpdateDopust:output_type -> dopust.Dopust
	7,  // 13: dopust.DopustService.DeleteDopust:output_type -> dopust.DeleteDopustResponse
	1,  // 14: dopust.DopustService.CreateZaposlen:output_type -> dopust.Zaposlen
	10, // 15: dopust.DopustService.GetZaposleni:output_type -> dopust.ZaposlenList
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_dopust_proto_init() }
func file_dopust_proto_init() {
	if File_dopust_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dopust_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dopust); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zaposlen); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZaposlen); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZaposlenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDopustRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDopustRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDopustRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDopustResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDopustiRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZaposleniParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZaposlenList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDopustiParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dopust_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DopustList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_dopust_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dopust_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dopust_proto_goTypes,
		DependencyIndexes: file_dopust_proto_depIdxs,
		MessageInfos:      file_dopust_proto_msgTypes,
	}.Build()
	File_dopust_proto = out.File
	file_dopust_proto_rawDesc = nil
	file_dopust_proto_goTypes = nil
	file_dopust_proto_depIdxs = nil
}