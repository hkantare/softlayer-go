/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package services

import (
	"github.ibm.com/riethm/gopherlayer/datatypes"
	"github.ibm.com/riethm/gopherlayer/session"
	"github.ibm.com/riethm/gopherlayer/sl"
)

// Every piece of hardware and network connection owned by SoftLayer is tracked physically by location and stored in the SoftLayer_Location data type. SoftLayer locations exist in parent/child relationships, a convenient way to track equipment from it's city, datacenter, server room, rack, then slot. Network backbones are tied to datacenters only, not to a room, rack, or slot.
type Location struct {
	Session *session.Session
	sl.Options
}

func GetLocationService(sess *session.Session) Location {
	return Location{Session: sess}
}

// Object Storage is only available in select datacenters. This method will return all the datacenters where object storage is available.
func (r *Location) GetAvailableObjectStorageDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location) GetBackboneDependents() (resp []datatypes.Network_Backbone_Location_Dependent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all datacenter locations. SoftLayer's datacenters exist in various cities and each contain one or more server rooms which house network and server infrastructure.
func (r *Location) GetDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location) GetDatacentersWithVirtualImageStoreServiceResourceRecord() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 or more groups. This will show which groups to which a location belongs.
func (r *Location) GetGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location) GetHardwareFirewalls() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location's physical address.
func (r *Location) GetLocationAddress() (resp datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location's Dedicated Rack member
func (r *Location) GetLocationReservationMember() (resp datatypes.Location_Reservation_Rack_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current locations status.
func (r *Location) GetLocationStatus() (resp datatypes.Location_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location) GetNetworkConfigurationAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location) GetObject() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of users online using SoftLayer's PPTP VPN service for a location.
func (r *Location) GetOnlinePptpVpnUserCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of users online using SoftLayer's SSL VPN service for a location.
func (r *Location) GetOnlineSslVpnUserCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location) GetPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 or more Price Groups. This will show which groups to which a location belongs.
func (r *Location) GetPriceGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 or more regions. This will show which regions to which a location belongs.
func (r *Location) GetRegions() (resp []datatypes.Location_Region, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location) GetTimezone() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 Bandwidth Pooling Group. This will show which group to which a location belongs.
func (r *Location) GetVdrGroup() (resp datatypes.Location_Group_Location_CrossReference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all datacenter locations. SoftLayer's datacenters exist in various cities and each contain one or more server rooms which house network and server infrastructure.
func (r *Location) GetViewableDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all viewable pop and datacenter locations.
func (r *Location) GetViewablePopsAndDataCenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all viewable network locations.
func (r *Location) GetViewablepointOfPresence() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all point of presence locations.
func (r *Location) GetpointOfPresence() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// SoftLayer_Location_Datacenter extends the [[SoftLayer_Location]] data type to include datacenter-specific properties.
type Location_Datacenter struct {
	Session *session.Session
	sl.Options
}

func GetLocationDatacenterService(sess *session.Session) Location_Datacenter {
	return Location_Datacenter{Session: sess}
}

// Retrieve
func (r *Location_Datacenter) GetActiveItemPresaleEvents() (resp []datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetActivePresaleEvents() (resp []datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Object Storage is only available in select datacenters. This method will return all the datacenters where object storage is available.
func (r *Location_Datacenter) GetAvailableObjectStorageDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetBackboneDependents() (resp []datatypes.Network_Backbone_Location_Dependent, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetBackendHardwareRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Subnets which are directly bound to one or more routers in a given datacenter, and currently allow routing.
func (r *Location_Datacenter) GetBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve This references relationship between brands, locations and countries associated with a user's account that are ineligible when ordering products. For example, the India datacenter may not be available on this brand for customers that live in Great Britain.
func (r *Location_Datacenter) GetBrandCountryRestrictions() (resp []datatypes.Brand_Restriction_Location_CustomerCountry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all datacenter locations. SoftLayer's datacenters exist in various cities and each contain one or more server rooms which house network and server infrastructure.
func (r *Location_Datacenter) GetDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Datacenter) GetDatacentersWithVirtualImageStoreServiceResourceRecord() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetFrontendHardwareRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 or more groups. This will show which groups to which a location belongs.
func (r *Location_Datacenter) GetGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetHardwareFirewalls() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetHardwareRouters() (resp []datatypes.Hardware, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location's physical address.
func (r *Location_Datacenter) GetLocationAddress() (resp datatypes.Account_Address, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location's Dedicated Rack member
func (r *Location_Datacenter) GetLocationReservationMember() (resp datatypes.Location_Reservation_Rack_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The current locations status.
func (r *Location_Datacenter) GetLocationStatus() (resp datatypes.Location_Status, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetNetworkConfigurationAttribute() (resp datatypes.Hardware_Attribute, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Datacenter) GetObject() (resp datatypes.Location_Datacenter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of users online using SoftLayer's PPTP VPN service for a location.
func (r *Location_Datacenter) GetOnlinePptpVpnUserCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The total number of users online using SoftLayer's SSL VPN service for a location.
func (r *Location_Datacenter) GetOnlineSslVpnUserCount() (resp int, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetPathString() (resp string, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetPresaleEvents() (resp []datatypes.Sales_Presale_Event, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 or more Price Groups. This will show which groups to which a location belongs.
func (r *Location_Datacenter) GetPriceGroups() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The regional group this datacenter belongs to.
func (r *Location_Datacenter) GetRegionalGroup() (resp datatypes.Location_Group_Regional, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetRegionalInternetRegistry() (resp datatypes.Network_Regional_Internet_Registry, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 or more regions. This will show which regions to which a location belongs.
func (r *Location_Datacenter) GetRegions() (resp []datatypes.Location_Region, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Retrieve all subnets that are eligible to be routed; those which the account has permission to associate with a vlan.
func (r *Location_Datacenter) GetRoutableBoundSubnets() (resp []datatypes.Network_Subnet, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve a graph of a SoftLayer datacenter's last 48 hours of network activity. Statistics graphs show traffic outbound from a datacenter on top and inbound traffic on the bottom followed by a legend of the network services tracked in the graph. getStatisticsGraphImage returns a PNG image of variable width and height depending on the number of services reported in the image.
func (r *Location_Datacenter) GetStatisticsGraphImage() (resp []byte, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Datacenter) GetTimezone() (resp datatypes.Locale_Timezone, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve A location can be a member of 1 Bandwidth Pooling Group. This will show which group to which a location belongs.
func (r *Location_Datacenter) GetVdrGroup() (resp datatypes.Location_Group_Location_CrossReference, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all datacenter locations. SoftLayer's datacenters exist in various cities and each contain one or more server rooms which house network and server infrastructure.
func (r *Location_Datacenter) GetViewableDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all viewable pop and datacenter locations.
func (r *Location_Datacenter) GetViewablePopsAndDataCenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all viewable network locations.
func (r *Location_Datacenter) GetViewablepointOfPresence() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve all point of presence locations.
func (r *Location_Datacenter) GetpointOfPresence() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Location_Group struct {
	Session *session.Session
	sl.Options
}

func GetLocationGroupService(sess *session.Session) Location_Group {
	return Location_Group{Session: sess}
}

//
func (r *Location_Group) GetAllObjects() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type for this location group.
func (r *Location_Group) GetLocationGroupType() (resp datatypes.Location_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The locations in a group.
func (r *Location_Group) GetLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Group) GetObject() (resp datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Location_Group_Pricing struct {
	Session *session.Session
	sl.Options
}

func GetLocationGroupPricingService(sess *session.Session) Location_Group_Pricing {
	return Location_Group_Pricing{Session: sess}
}

//
func (r *Location_Group_Pricing) GetAllObjects() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type for this location group.
func (r *Location_Group_Pricing) GetLocationGroupType() (resp datatypes.Location_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The locations in a group.
func (r *Location_Group_Pricing) GetLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Group_Pricing) GetObject() (resp datatypes.Location_Group_Pricing, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The prices that this pricing location group limits. All of these prices will only be available in the locations defined by this pricing location group.
func (r *Location_Group_Pricing) GetPrices() (resp []datatypes.Product_Item_Price, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Location_Group_Regional struct {
	Session *session.Session
	sl.Options
}

func GetLocationGroupRegionalService(sess *session.Session) Location_Group_Regional {
	return Location_Group_Regional{Session: sess}
}

//
func (r *Location_Group_Regional) GetAllObjects() (resp []datatypes.Location_Group, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The datacenters in a group.
func (r *Location_Group_Regional) GetDatacenters() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The type for this location group.
func (r *Location_Group_Regional) GetLocationGroupType() (resp datatypes.Location_Group_Type, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The locations in a group.
func (r *Location_Group_Regional) GetLocations() (resp []datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Group_Regional) GetObject() (resp datatypes.Location_Group_Regional, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The preferred datacenters of a group.
func (r *Location_Group_Regional) GetPreferredDatacenter() (resp datatypes.Location_Datacenter, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Location_Reservation struct {
	Session *session.Session
	sl.Options
}

func GetLocationReservationService(sess *session.Session) Location_Reservation {
	return Location_Reservation{Session: sess}
}

// Retrieve The account that a billing item belongs to.
func (r *Location_Reservation) GetAccount() (resp datatypes.Account, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Reservation) GetAccountReservations() (resp []datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The bandwidth allotment that the reservation belongs to.
func (r *Location_Reservation) GetAllotment() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The bandwidth allotment that the reservation belongs to.
func (r *Location_Reservation) GetBillingItem() (resp datatypes.Billing_Item, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve The datacenter location that the reservation belongs to.
func (r *Location_Reservation) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Rack information for the reservation
func (r *Location_Reservation) GetLocationReservationRack() (resp datatypes.Location_Reservation_Rack, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Reservation) GetObject() (resp datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Location_Reservation_Rack struct {
	Session *session.Session
	sl.Options
}

func GetLocationReservationRackService(sess *session.Session) Location_Reservation_Rack {
	return Location_Reservation_Rack{Session: sess}
}

// Retrieve The bandwidth allotment that the reservation belongs to.
func (r *Location_Reservation_Rack) GetAllotment() (resp datatypes.Network_Bandwidth_Version1_Allotment, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve Members of the rack.
func (r *Location_Reservation_Rack) GetChildren() (resp []datatypes.Location_Reservation_Rack_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Reservation_Rack) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Reservation_Rack) GetLocationReservation() (resp datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Reservation_Rack) GetObject() (resp datatypes.Location_Reservation_Rack, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
type Location_Reservation_Rack_Member struct {
	Session *session.Session
	sl.Options
}

func GetLocationReservationRackMemberService(sess *session.Session) Location_Reservation_Rack_Member {
	return Location_Reservation_Rack_Member{Session: sess}
}

// Retrieve Location relation for the rack member
func (r *Location_Reservation_Rack_Member) GetLocation() (resp datatypes.Location, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

// Retrieve
func (r *Location_Reservation_Rack_Member) GetLocationReservationRack() (resp datatypes.Location_Reservation, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}

//
func (r *Location_Reservation_Rack_Member) GetObject() (resp datatypes.Location_Reservation_Rack_Member, err error) {
	err = invokeMethod(nil, r.Session, &r.Options, &resp)
	return
}
