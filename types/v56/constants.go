package types

const (
	// NsOvf the ovf xml namespace url
	NsOvf = "http://schemas.dmtf.org/ovf/envelope/1"
	// NsXMLSchema the xml schema namespace url
	NsXMLSchema = "http://www.w3.org/2001/XMLSchema-instance"
	// NsVCloud vcloud xml namespace url
	NsVCloud = "http://www.vmware.com/vcloud/v1.5"
)

const (
	// Empty Mime
	MimeEmpty = ""
	// MimeOrgList mime for org list
	MimeOrgList = "application/vnd.vmware.vcloud.orgList+xml"
	// MimeOrg mime for org
	MimeOrg = "application/vnd.vmware.vcloud.org+xml"
	// MimeAdminOrg mime for adminOrg
	MimeAdminOrg = "application/vnd.vmware.admin.organization+xml"
	// MimeCatalog mime for catalog
	MimeCatalog = "application/vnd.vmware.vcloud.catalog+xml"
	// MimeAdminCatalog mime for a adminCatalog
	MimeAdminCatalog = "application/vnd.vmware.admin.catalog+xml"
	// MimeCatalogItem mime for catalog item
	MimeCatalogItem = "application/vnd.vmware.vcloud.catalogItem+xml"
	// MimeVDC mime for a VDC
	MimeVDC = "application/vnd.vmware.vcloud.vdc+xml"
	// MimeVAppTemplate mime for a vapp template
	MimeVAppTemplate = "application/vnd.vmware.vcloud.vAppTemplate+xml"
	// MimeInstantiateVAppTemplate mime fore instantiate VApp template params
	MimeInstantiateVAppTemplate = "application/vnd.vmware.vcloud.instantiateVAppTemplateParams+xml"
	// MimeVApp mime for a vApp
	MimeVApp = "application/vnd.vmware.vcloud.vApp+xml"
	// MimeQueryRecords mime for the query records
	MimeQueryRecords = "application/vnd.vmware.vchs.query.records+xml"
	// MimeAPIExtensibility mime for api extensibility
	MimeAPIExtensibility = "application/vnd.vmware.vcloud.apiextensibility+xml"
	// MimeEntity mime for vcloud entity
	MimeEntity = "application/vnd.vmware.vcloud.entity+xml"
	// MimeQueryList mime for query list
	MimeQueryList = "application/vnd.vmware.vcloud.query.queryList+xml"
	// MimeSession mime for a session
	MimeSession = "application/vnd.vmware.vcloud.session+xml"
	// MimeTask mime for task
	MimeTask = "application/vnd.vmware.vcloud.task+xml"
	// MimeError mime for error
	MimeError = "application/vnd.vmware.vcloud.error+xml"
	// MimeNetwork mime for a network
	MimeNetwork = "application/vnd.vmware.vcloud.network+xml"
	// MimeMedia mime for Media object
	MimeMedia             = "application/vnd.vmware.vcloud.media+xml"
	MimeComposeVAppParams = "application/vnd.vmware.vcloud.composeVAppParams+xml"
	//MimeDiskCreateParams mime for create independent disk
	MimeDiskCreateParams = "application/vnd.vmware.vcloud.diskCreateParams+xml"
	// Mime for attach or detach independent disk
	MimeDiskAttachOrDetachParams = "application/vnd.vmware.vcloud.diskAttachOrDetachParams+xml"
	// Mime for Disk
	MimeDisk = "application/vnd.vmware.vcloud.disk+xml"
	// Mime for VMs
	MimeVMs = "application/vnd.vmware.vcloud.vms+xml"
)

const (
	// HTTPGet the http GET method
	HTTPGet = "GET"
	// HTTPPost the http POST method
	HTTPPost = "POST"
	// HTTPPut the http PUT method
	HTTPPut = "PUT"
	// HTTPPatch the http PATCH method
	HTTPPatch = "PATCH"
	// HTTPDelete the http DELETE method
	HTTPDelete = "DELETE"
)
