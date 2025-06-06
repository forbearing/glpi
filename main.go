package main

import (
	model_admin "glpi/model/admin"
	model_asset "glpi/model/asset"
	model_assist "glpi/model/assistance"
	model_mgmt "glpi/model/management"
	model_setup "glpi/model/setup"
	model_comp "glpi/model/setup/component"
	model_dropdown_appl "glpi/model/setup/dropdown/appliance"
	model_dropdown_assist "glpi/model/setup/dropdown/assistance"
	model_dropdown_black "glpi/model/setup/dropdown/blacklist"
	model_dropdown_cable "glpi/model/setup/dropdown/cable"
	model_dropdown_calen "glpi/model/setup/dropdown/calendar"
	model_dropdown_camera "glpi/model/setup/dropdown/camera"
	model_dropdown_common "glpi/model/setup/dropdown/common"
	model_dropdown_inet "glpi/model/setup/dropdown/internet"
	model_dropdown_ldap "glpi/model/setup/dropdown/ldap"
	model_dropdown_mgmt "glpi/model/setup/dropdown/mangement"
	model_dropdown_model "glpi/model/setup/dropdown/model"
	model_dropdown_net "glpi/model/setup/dropdown/network"
	model_dropdown_os "glpi/model/setup/dropdown/os"
	model_dropdown_other "glpi/model/setup/dropdown/others"
	model_dropdown_power "glpi/model/setup/dropdown/power"
	model_dropdown_soft "glpi/model/setup/dropdown/software"
	model_dropdown_sso "glpi/model/setup/dropdown/sso"
	model_dropdown_tool "glpi/model/setup/dropdown/tool"
	model_dropdown_type "glpi/model/setup/dropdown/type"
	model_dropdown_user "glpi/model/setup/dropdown/user"
	model_dropdown_vm "glpi/model/setup/dropdown/vm"
	model_tool "glpi/model/tool"

	"github.com/forbearing/golib/bootstrap"
	"github.com/forbearing/golib/router"
	"github.com/forbearing/golib/util"
)

func main() {
	util.RunOrDie(bootstrap.Bootstrap)

	asset := router.API().Group("/asset")
	router.Register[*model_asset.Computer](asset, "/computer")
	router.Register[*model_asset.Monitor](asset, "/monitor")
	router.Register[*model_asset.Software](asset, "/software")
	router.Register[*model_asset.NetworkEquipment](asset, "/network_equipment")
	router.Register[*model_asset.Peripheral](asset, "/peripheral")
	router.Register[*model_asset.Printer](asset, "/printer")
	router.Register[*model_asset.Cartridge](asset, "/cartridge")
	router.Register[*model_asset.Consumable](asset, "/consumable")
	router.Register[*model_asset.Phone](asset, "/phone")
	router.Register[*model_asset.Rack](asset, "/rack")
	router.Register[*model_asset.Enclosure](asset, "/enclosure")
	router.Register[*model_asset.PDU](asset, "/pdu")
	router.Register[*model_asset.PassiveDeviceEquipment](asset, "/passive_device_equipment")
	router.Register[*model_asset.Unmanaged](asset, "/unmanaged")
	router.Register[*model_asset.Cable](asset, "/cable")
	router.Register[*model_asset.DeviceSimcard](asset, "/device_simcard")

	assistance := router.API().Group("/assistance")
	router.Register[*model_assist.Ticket](assistance, "/ticket")
	router.Register[*model_assist.Problem](assistance, "/problem")
	router.Register[*model_assist.Change](assistance, "/change")
	router.Register[*model_assist.PlanningExternalEvent](assistance, "/planning_external_event")
	router.Register[*model_assist.TicketRecurrent](assistance, "/ticket_recurrent")
	router.Register[*model_assist.ChangeRecurrent](assistance, "/change_recurrent")

	management := router.API().Group("/management")
	router.Register[*model_mgmt.SoftwareLicense](management, "/software_license")
	router.Register[*model_mgmt.Budget](management, "/budget")
	router.Register[*model_mgmt.Supplier](management, "/supplier")
	router.Register[*model_mgmt.Contact](management, "/contact")
	router.Register[*model_mgmt.Contract](management, "/contract")
	router.Register[*model_mgmt.Document](management, "/document")
	router.Register[*model_mgmt.Line](management, "/line")
	router.Register[*model_mgmt.Certificate](management, "/certificate")
	router.Register[*model_mgmt.Datacenter](management, "/datacenter")
	router.Register[*model_mgmt.Cluster](management, "/cluster")
	router.Register[*model_mgmt.Domain](management, "/domain")
	router.Register[*model_mgmt.Appliance](management, "/appliance")
	router.Register[*model_mgmt.Database](management, "/database")

	tool := router.API().Group("/tool")
	router.Register[*model_tool.Project](tool, "/project")
	router.Register[*model_tool.Reminder](tool, "/reminder")
	router.Register[*model_tool.RSSFeed](tool, "/rss_feed")
	router.Register[*model_tool.Knowbase](tool, "/knowbase")
	router.Register[*model_tool.Reservation](tool, "/reservation")
	router.Register[*model_tool.SavedSearch](tool, "/saved_search")

	admin := router.API().Group("/admin")
	router.Register[*model_admin.User](admin, "/user")
	router.Register[*model_admin.Group](admin, "/group")
	router.Register[*model_admin.Entity](admin, "/entity")
	router.Register[*model_admin.Rule](admin, "/rule")
	router.Register[*model_admin.Directionary](admin, "/directionary")
	router.Register[*model_admin.Profile](admin, "/profile")
	router.Register[*model_admin.QueuedNotification](admin, "/queued_notification")
	router.Register[*model_admin.Log](admin, "/log")
	router.Register[*model_admin.Inventory](admin, "/inventory")

	setup := router.API().Group("/setup")
	dropdown := setup.Group("/dropdown")
	component := router.API().Group("/component")

	router.Register[*model_dropdown_common.Location](dropdown, "/location")
	router.Register[*model_dropdown_common.Status](dropdown, "/status")
	router.Register[*model_dropdown_common.Manufacturer](dropdown, "/manufacturer")
	router.Register[*model_dropdown_common.Blacklist](dropdown, "/blacklist")
	router.Register[*model_dropdown_common.BlacklistedMailContent](dropdown, "/blacklisted_mailcontent")

	router.Register[*model_dropdown_assist.ITILCategory](dropdown, "/itil_category")
	router.Register[*model_dropdown_assist.TaskCategory](dropdown, "/task_category")
	router.Register[*model_dropdown_assist.TaskTemplate](dropdown, "/task_template")
	router.Register[*model_dropdown_assist.SolutionType](dropdown, "/solution_type")
	router.Register[*model_dropdown_assist.SolutionTemplate](dropdown, "/solution_template")
	router.Register[*model_dropdown_assist.RequestType](dropdown, "/request_type")
	router.Register[*model_dropdown_assist.ITILFollowupTemplate](dropdown, "/itil_followup_template")
	router.Register[*model_dropdown_assist.ProjectStatus](dropdown, "/project_status")
	router.Register[*model_dropdown_assist.ProjectType](dropdown, "/project_type")
	router.Register[*model_dropdown_assist.ProjectTaskType](dropdown, "/project_task_type")
	router.Register[*model_dropdown_assist.ProjectTaskTemplate](dropdown, "/project_task_template")
	router.Register[*model_dropdown_assist.PlanningExternalEventTemplate](dropdown, "/planning_external_event_template")
	router.Register[*model_dropdown_assist.PlanningEventCategory](dropdown, "/planning_event_category")
	router.Register[*model_dropdown_assist.PendingReason](dropdown, "/pending_reason")

	router.Register[*model_dropdown_type.ComputerType](dropdown, "/computer_type")
	router.Register[*model_dropdown_type.NetworkEquipmentType](dropdown, "/network_equipment_type")
	router.Register[*model_dropdown_type.PrinterType](dropdown, "/printer_type")
	router.Register[*model_dropdown_type.MonitorType](dropdown, "/monitor_type")
	router.Register[*model_dropdown_type.PeripheralType](dropdown, "/peripheral_type")
	router.Register[*model_dropdown_type.PhoneType](dropdown, "/phone_type")
	router.Register[*model_dropdown_type.SoftwareLicenseType](dropdown, "/software_license_type")
	router.Register[*model_dropdown_type.CartridgeType](dropdown, "/cartridge_type")
	router.Register[*model_dropdown_type.ConsumableType](dropdown, "/consumable_type")
	router.Register[*model_dropdown_type.ContractType](dropdown, "/contract_type")
	router.Register[*model_dropdown_type.ContactType](dropdown, "/contact_type")
	router.Register[*model_dropdown_type.DeviceGenericType](dropdown, "/device_generic_type")
	router.Register[*model_dropdown_type.DeviceSensorType](dropdown, "/device_sensor_type")
	router.Register[*model_dropdown_type.DeviceMemoryType](dropdown, "/device_memory_type")
	router.Register[*model_dropdown_type.DeviceCaseType](dropdown, "/device_case_type")
	router.Register[*model_dropdown_type.DeviceSimcardType](dropdown, "/device_simcard_type")
	router.Register[*model_dropdown_type.SupplierType](dropdown, "/supplier_type")
	router.Register[*model_dropdown_type.InterfaceType](dropdown, "/interface_type")
	router.Register[*model_dropdown_type.PhonePowerSupplyType](dropdown, "/phone_power_supply_type")
	router.Register[*model_dropdown_type.FilesystemType](dropdown, "/filesystem_type")
	router.Register[*model_dropdown_type.CertificateType](dropdown, "/certificate_type")
	router.Register[*model_dropdown_type.BudgetType](dropdown, "/budget_type")
	router.Register[*model_dropdown_type.LineType](dropdown, "/line_type")
	router.Register[*model_dropdown_type.RackType](dropdown, "/rack_type")
	router.Register[*model_dropdown_type.PDUType](dropdown, "/pdu_type")
	router.Register[*model_dropdown_type.PassiveDeviceEquipmentType](dropdown, "/passive_device_equipment_type")
	router.Register[*model_dropdown_type.ClusterType](dropdown, "/cluster_type")
	router.Register[*model_dropdown_type.DatabaseInstanceType](dropdown, "/database_instance_type")
	router.Register[*model_dropdown_model.ComputerModel](dropdown, "/computer_model")
	router.Register[*model_dropdown_model.NetworkEquipmentModel](dropdown, "/network_equipment_model")
	router.Register[*model_dropdown_model.PrinterModel](dropdown, "/printer_model")
	router.Register[*model_dropdown_model.MonitorModel](dropdown, "/monitor_model")
	router.Register[*model_dropdown_model.PeripheralModel](dropdown, "/peripheral_model")
	router.Register[*model_dropdown_model.PhoneModel](dropdown, "/phone_model")
	router.Register[*model_dropdown_model.DeviceCameraModel](dropdown, "/device_camera_model")
	router.Register[*model_dropdown_model.DeviceCaseModel](dropdown, "/device_case_model")
	router.Register[*model_dropdown_model.DeviceControlModel](dropdown, "/device_control_model")
	router.Register[*model_dropdown_model.DeviceDriveModel](dropdown, "/device_drive_model")
	router.Register[*model_dropdown_model.DeviceGenericModel](dropdown, "/device_generic_model")
	router.Register[*model_dropdown_model.DeviceGraphicCardModel](dropdown, "/device_graphic_card_model")
	router.Register[*model_dropdown_model.DeviceHardDriveModel](dropdown, "/device_hard_drive_model")
	router.Register[*model_dropdown_model.DeviceMemoryModel](dropdown, "/device_memory_model")
	router.Register[*model_dropdown_model.DeviceMotherBoardModel](dropdown, "/device_mother_board_model")
	router.Register[*model_dropdown_model.DeviceNetworkCardModel](dropdown, "/device_network_card_model")
	router.Register[*model_dropdown_model.DevicePciModel](dropdown, "/device_pci_model")
	router.Register[*model_dropdown_model.DevicePowerSupplyModel](dropdown, "/device_power_supply_model")
	router.Register[*model_dropdown_model.DeviceProcessorModel](dropdown, "/device_processor_model")
	router.Register[*model_dropdown_model.DeviceSoundCardModel](dropdown, "/device_sound_card_model")
	router.Register[*model_dropdown_model.DeviceSensorModel](dropdown, "/device_sensor_model")
	router.Register[*model_dropdown_model.RackModel](dropdown, "/rack_model")
	router.Register[*model_dropdown_model.EnclosureModel](dropdown, "/enclosure_model")
	router.Register[*model_dropdown_model.PDUModel](dropdown, "/pdu_model")
	router.Register[*model_dropdown_model.PassiveDeviceEquipmentModel](dropdown, "/passive_device_equipment_model")
	router.Register[*model_dropdown_vm.VirtualMachineType](dropdown, "/virtual_machine_type")
	router.Register[*model_dropdown_vm.VirtualMachineStatus](dropdown, "/virtual_machine_status")
	router.Register[*model_dropdown_vm.VirtualMachineSystem](dropdown, "/virtual_machine_system")
	router.Register[*model_dropdown_mgmt.DocumentCategory](dropdown, "/document_category")
	router.Register[*model_dropdown_mgmt.DocumentType](dropdown, "/document_type")
	router.Register[*model_dropdown_mgmt.BusinessCriticity](dropdown, "/business_criticity")
	router.Register[*model_dropdown_mgmt.DatabaseInstanceCategory](dropdown, "/database_instance_category")
	router.Register[*model_dropdown_tool.KnowbaseCategory](dropdown, "/knowbase_category")
	router.Register[*model_dropdown_calen.Calendar](dropdown, "/calendar")
	router.Register[*model_dropdown_calen.Holiday](dropdown, "/holiday")
	router.Register[*model_dropdown_os.OperatingSystem](dropdown, "/operating_system")
	router.Register[*model_dropdown_os.OperatingSystemVersion](dropdown, "/operating_system_version")
	router.Register[*model_dropdown_os.OperatingSystemServicePack](dropdown, "/operating_system_service_pack")
	router.Register[*model_dropdown_os.OperatingSystemArchitecture](dropdown, "/operating_system_architecture")
	router.Register[*model_dropdown_os.OperatingSystemEdition](dropdown, "/operating_system_edition")
	router.Register[*model_dropdown_os.OperatingSystemKernel](dropdown, "/operating_system_kernel")
	router.Register[*model_dropdown_os.OperatingSystemKernelVersion](dropdown, "/operating_system_kernel_version")
	router.Register[*model_dropdown_os.AutoUpdateSystem](dropdown, "/auto_update_system")
	router.Register[*model_dropdown_net.NetworkInterface](dropdown, "/network_interface")
	router.Register[*model_dropdown_net.Network](dropdown, "/network")
	router.Register[*model_dropdown_net.NetworkPortType](dropdown, "/network_port_type")
	router.Register[*model_dropdown_net.Vlan](dropdown, "/vlan")
	router.Register[*model_dropdown_net.LineOperator](dropdown, "/line_operator")
	router.Register[*model_dropdown_net.DomainType](dropdown, "/domain_type")
	router.Register[*model_dropdown_net.DomainRelation](dropdown, "/domain_relation")
	router.Register[*model_dropdown_net.DomainRecordType](dropdown, "/domain_record_type")
	router.Register[*model_dropdown_net.NetworkPortFiberChannelType](dropdown, "/network_port_fiber_channel_type")
	router.Register[*model_dropdown_cable.CableType](dropdown, "/cable_type")
	router.Register[*model_dropdown_cable.CableStrand](dropdown, "/cable_strand")
	router.Register[*model_dropdown_cable.SocketModel](dropdown, "/socket_model")
	router.Register[*model_dropdown_inet.IPNetwork](dropdown, "/ip_network")
	router.Register[*model_dropdown_inet.FQDN](dropdown, "/fqdn")
	router.Register[*model_dropdown_inet.WifiNetwork](dropdown, "/wifi_network")
	router.Register[*model_dropdown_inet.NetworkName](dropdown, "/network_name")
	router.Register[*model_dropdown_soft.SoftwareCategory](dropdown, "/software_category")
	router.Register[*model_dropdown_user.UserTitle](dropdown, "/user_title")
	router.Register[*model_dropdown_user.UserCategory](dropdown, "/user_category")
	router.Register[*model_dropdown_ldap.RuleRightParameter](dropdown, "/rule_right_parameter")
	router.Register[*model_dropdown_sso.SSOVariable](dropdown, "/sso_variable")
	router.Register[*model_dropdown_black.FieldBlacklist](dropdown, "/field_blacklist")
	router.Register[*model_dropdown_power.Plug](dropdown, "/plug")
	router.Register[*model_dropdown_appl.ApplianceType](dropdown, "/appliance_type")
	router.Register[*model_dropdown_appl.ApplianceEnvironment](dropdown, "/appliance_environment")
	router.Register[*model_dropdown_camera.ImageFormat](dropdown, "/image_format")
	router.Register[*model_dropdown_other.USBVendor](dropdown, "/usb_vendor")
	router.Register[*model_dropdown_other.PCIVendor](dropdown, "/pci_vendor")

	router.Register[*model_comp.DeviceBattery](component, "/device_battery")
	router.Register[*model_comp.DeviceCamera](component, "/device_camera")
	router.Register[*model_comp.DeviceCase](component, "/device_case")
	router.Register[*model_comp.DeviceControl](component, "/device_control")
	router.Register[*model_comp.DeviceDrive](component, "/device_drive")
	router.Register[*model_comp.DeviceFirmware](component, "/device_firmware")
	router.Register[*model_comp.DeviceGeneric](component, "/device_generic")
	router.Register[*model_comp.DeviceGraphicCard](component, "/device_graphic_card")
	router.Register[*model_comp.DeviceHardDrive](component, "/device_hard_drive")
	router.Register[*model_comp.DeviceMemory](component, "/device_memory")
	router.Register[*model_comp.DeviceNetworkCard](component, "/device_network_card")
	router.Register[*model_comp.DevicePCI](component, "/device_pci")
	router.Register[*model_comp.DevicePowerSupply](component, "/device_power_supply")
	router.Register[*model_comp.DeviceProcessor](component, "/device_processor")
	router.Register[*model_comp.DeviceSensor](component, "/device_sensor")
	router.Register[*model_comp.DeviceSimcard](component, "/device_simcard")
	router.Register[*model_comp.DeviceSoundCard](component, "/device_sound_card")
	router.Register[*model_comp.DeviceMotherboard](component, "/device_motherboard")

	router.Register[*model_setup.SLM](setup, "/slm")
	router.Register[*model_setup.SLA](setup, "/sla")
	router.Register[*model_setup.OLA](setup, "/ola")
	router.Register[*model_setup.Config](setup, "/config")
	router.Register[*model_setup.FieldUnicity](setup, "/field_unicity")
	router.Register[*model_setup.CronTask](setup, "/cron_task")
	router.Register[*model_setup.CronTaskLog](setup, "/cron_task_log")
	router.Register[*model_setup.AuthLDAP](setup, "/auth_ldap")
	router.Register[*model_setup.AuthMail](setup, "/auth_mail")
	router.Register[*model_setup.MailCollector](setup, "/mail_collector")
	router.Register[*model_setup.Link](setup, "/link")

	util.RunOrDie(router.Run)
}
