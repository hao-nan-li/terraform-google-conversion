// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//	This file is automatically generated by Magic Modules and manual
//	changes will be clobbered when the file is regenerated.
//
//	Please read more about how to change this file in
//	.github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

// ResourceConverter returns a map of terraform resource types (i.e. `google_project`)
// to a slice of ResourceConverters.
//
// Modelling of relationships:
// terraform resources to CAI assets as []ResourceConverter:
// 1:1 = [ResourceConverter{Convert: convertAbc}]                  (len=1)
// 1:N = [ResourceConverter{Convert: convertAbc}, ...]             (len=N)
// N:1 = [ResourceConverter{Convert: convertAbc, merge: mergeAbc}] (len=1)
func ResourceConverters() map[string][]ResourceConverter {
	return map[string][]ResourceConverter{
		"google_compute_address":                                  {resourceConverterComputeAddress()},
		"google_compute_firewall":                                 {resourceConverterComputeFirewall()},
		"google_compute_disk":                                     {resourceConverterComputeDisk()},
		"google_compute_forwarding_rule":                          {resourceConverterComputeForwardingRule()},
		"google_compute_global_address":                           {resourceConverterComputeGlobalAddress()},
		"google_compute_global_forwarding_rule":                   {resourceConverterComputeGlobalForwardingRule()},
		"google_compute_instance":                                 {resourceConverterComputeInstance()},
		"google_compute_network":                                  {resourceConverterComputeNetwork()},
		"google_compute_security_policy":                          {resourceConverterComputeSecurityPolicy()},
		"google_compute_snapshot":                                 {resourceConverterComputeSnapshot()},
		"google_compute_subnetwork":                               {resourceConverterComputeSubnetwork()},
		"google_compute_ssl_policy":                               {resourceConverterComputeSslPolicy()},
		"google_compute_target_ssl_proxy":                         {resourceConverterComputeTargetSslProxy()},
		"google_dns_managed_zone":                                 {resourceConverterDNSManagedZone()},
		"google_dns_policy":                                       {resourceConverterDNSPolicy()},
		"google_storage_bucket":                                   {resourceConverterStorageBucket()},
		"google_sql_database_instance":                            {resourceConverterSQLDatabaseInstance()},
		"google_sql_database":                                     {resourceConverterSQLDatabase()},
		"google_container_cluster":                                {resourceConverterContainerCluster()},
		"google_container_node_pool":                              {resourceConverterContainerNodePool()},
		"google_bigquery_dataset":                                 {resourceConverterBigQueryDataset()},
		"google_bigquery_dataset_iam_policy":                      {resourceConverterBigqueryDatasetIamPolicy()},
		"google_bigquery_dataset_iam_binding":                     {resourceConverterBigqueryDatasetIamBinding()},
		"google_bigquery_dataset_iam_member":                      {resourceConverterBigqueryDatasetIamMember()},
		"google_bigquery_table":                                   {resourceConverterBigQueryTable()},
		"google_redis_instance":                                   {resourceConverterRedisInstance()},
		"google_spanner_database":                                 {resourceConverterSpannerDatabase()},
		"google_spanner_database_iam_policy":                      {resourceConverterSpannerDatabaseIamPolicy()},
		"google_spanner_database_iam_binding":                     {resourceConverterSpannerDatabaseIamBinding()},
		"google_spanner_database_iam_member":                      {resourceConverterSpannerDatabaseIamMember()},
		"google_spanner_instance":                                 {resourceConverterSpannerInstance()},
		"google_spanner_instance_iam_policy":                      {resourceConverterSpannerInstanceIamPolicy()},
		"google_spanner_instance_iam_binding":                     {resourceConverterSpannerInstanceIamBinding()},
		"google_spanner_instance_iam_member":                      {resourceConverterSpannerInstanceIamMember()},
		"google_project_service":                                  {resourceConverterServiceUsage()},
		"google_pubsub_lite_reservation":                          {resourceConverterPubsubLiteReservation()},
		"google_pubsub_lite_subscription":                         {resourceConverterPubsubLiteSubscription()},
		"google_pubsub_lite_topic":                                {resourceConverterPubsubLiteTopic()},
		"google_pubsub_schema":                                    {resourceConverterPubsubSchema()},
		"google_pubsub_subscription":                              {resourceConverterPubsubSubscription()},
		"google_pubsub_subscription_iam_policy":                   {resourceConverterPubsubSubscriptionIamPolicy()},
		"google_pubsub_subscription_iam_binding":                  {resourceConverterPubsubSubscriptionIamBinding()},
		"google_pubsub_subscription_iam_member":                   {resourceConverterPubsubSubscriptionIamMember()},
		"google_storage_bucket_iam_policy":                        {resourceConverterStorageBucketIamPolicy()},
		"google_storage_bucket_iam_binding":                       {resourceConverterStorageBucketIamBinding()},
		"google_storage_bucket_iam_member":                        {resourceConverterStorageBucketIamMember()},
		"google_pubsub_topic":                                     {resourceConverterPubsubTopic()},
		"google_kms_crypto_key":                                   {resourceConverterKMSCryptoKey()},
		"google_kms_key_ring":                                     {resourceConverterKMSKeyRing()},
		"google_filestore_instance":                               {resourceConverterFilestoreInstance()},
		"google_access_context_manager_service_perimeter":         {resourceConverterAccessContextManagerServicePerimeter()},
		"google_access_context_manager_access_policy":             {resourceConverterAccessContextManagerAccessPolicy()},
		"google_cloud_run_service":                                {resourceConverterCloudRunService()},
		"google_cloud_run_domain_mapping":                         {resourceConverterCloudRunDomainMapping()},
		"google_cloudfunctions_function":                          {resourceConverterCloudFunctionsCloudFunction()},
		"google_monitoring_notification_channel":                  {resourceConverterMonitoringNotificationChannel()},
		"google_monitoring_alert_policy":                          {resourceConverterMonitoringAlertPolicy()},
		"google_access_context_manager_access_policy_iam_policy":  {resourceConverterAccessContextManagerAccessPolicyIamPolicy()},
		"google_access_context_manager_access_policy_iam_binding": {resourceConverterAccessContextManagerAccessPolicyIamBinding()},
		"google_access_context_manager_access_policy_iam_member":  {resourceConverterAccessContextManagerAccessPolicyIamMember()},
		"google_apigee_environment_iam_policy":                    {resourceConverterApigeeEnvironmentIamPolicy()},
		"google_apigee_environment_iam_binding":                   {resourceConverterApigeeEnvironmentIamBinding()},
		"google_apigee_environment_iam_member":                    {resourceConverterApigeeEnvironmentIamMember()},
		"google_artifact_registry_repository_iam_policy":          {resourceConverterArtifactRegistryRepositoryIamPolicy()},
		"google_artifact_registry_repository_iam_binding":         {resourceConverterArtifactRegistryRepositoryIamBinding()},
		"google_artifact_registry_repository_iam_member":          {resourceConverterArtifactRegistryRepositoryIamMember()},
		"google_bigquery_table_iam_policy":                        {resourceConverterBigQueryTableIamPolicy()},
		"google_bigquery_table_iam_binding":                       {resourceConverterBigQueryTableIamBinding()},
		"google_bigquery_table_iam_member":                        {resourceConverterBigQueryTableIamMember()},
		"google_bigquery_analytics_hub_data_exchange_iam_policy":  {resourceConverterBigqueryAnalyticsHubDataExchangeIamPolicy()},
		"google_bigquery_analytics_hub_data_exchange_iam_binding": {resourceConverterBigqueryAnalyticsHubDataExchangeIamBinding()},
		"google_bigquery_analytics_hub_data_exchange_iam_member":  {resourceConverterBigqueryAnalyticsHubDataExchangeIamMember()},
		"google_bigquery_analytics_hub_listing_iam_policy":        {resourceConverterBigqueryAnalyticsHubListingIamPolicy()},
		"google_bigquery_analytics_hub_listing_iam_binding":       {resourceConverterBigqueryAnalyticsHubListingIamBinding()},
		"google_bigquery_analytics_hub_listing_iam_member":        {resourceConverterBigqueryAnalyticsHubListingIamMember()},
		"google_bigquery_connection_iam_policy":                   {resourceConverterBigqueryConnectionConnectionIamPolicy()},
		"google_bigquery_connection_iam_binding":                  {resourceConverterBigqueryConnectionConnectionIamBinding()},
		"google_bigquery_connection_iam_member":                   {resourceConverterBigqueryConnectionConnectionIamMember()},
		"google_binary_authorization_attestor_iam_policy":         {resourceConverterBinaryAuthorizationAttestorIamPolicy()},
		"google_binary_authorization_attestor_iam_binding":        {resourceConverterBinaryAuthorizationAttestorIamBinding()},
		"google_binary_authorization_attestor_iam_member":         {resourceConverterBinaryAuthorizationAttestorIamMember()},
		"google_cloudfunctions_function_iam_policy":               {resourceConverterCloudFunctionsCloudFunctionIamPolicy()},
		"google_cloudfunctions_function_iam_binding":              {resourceConverterCloudFunctionsCloudFunctionIamBinding()},
		"google_cloudfunctions_function_iam_member":               {resourceConverterCloudFunctionsCloudFunctionIamMember()},
		"google_cloudfunctions2_function_iam_policy":              {resourceConverterCloudfunctions2functionIamPolicy()},
		"google_cloudfunctions2_function_iam_binding":             {resourceConverterCloudfunctions2functionIamBinding()},
		"google_cloudfunctions2_function_iam_member":              {resourceConverterCloudfunctions2functionIamMember()},
		"google_cloudiot_registry_iam_policy":                     {resourceConverterCloudIotDeviceRegistryIamPolicy()},
		"google_cloudiot_registry_iam_binding":                    {resourceConverterCloudIotDeviceRegistryIamBinding()},
		"google_cloudiot_registry_iam_member":                     {resourceConverterCloudIotDeviceRegistryIamMember()},
		"google_cloud_run_service_iam_policy":                     {resourceConverterCloudRunServiceIamPolicy()},
		"google_cloud_run_service_iam_binding":                    {resourceConverterCloudRunServiceIamBinding()},
		"google_cloud_run_service_iam_member":                     {resourceConverterCloudRunServiceIamMember()},
		"google_cloud_tasks_queue_iam_policy":                     {resourceConverterCloudTasksQueueIamPolicy()},
		"google_cloud_tasks_queue_iam_binding":                    {resourceConverterCloudTasksQueueIamBinding()},
		"google_cloud_tasks_queue_iam_member":                     {resourceConverterCloudTasksQueueIamMember()},
		"google_compute_backend_bucket_iam_policy":                {resourceConverterComputeBackendBucketIamPolicy()},
		"google_compute_backend_bucket_iam_binding":               {resourceConverterComputeBackendBucketIamBinding()},
		"google_compute_backend_bucket_iam_member":                {resourceConverterComputeBackendBucketIamMember()},
		"google_compute_backend_service_iam_policy":               {resourceConverterComputeBackendServiceIamPolicy()},
		"google_compute_backend_service_iam_binding":              {resourceConverterComputeBackendServiceIamBinding()},
		"google_compute_backend_service_iam_member":               {resourceConverterComputeBackendServiceIamMember()},
		"google_compute_region_backend_service_iam_policy":        {resourceConverterComputeRegionBackendServiceIamPolicy()},
		"google_compute_region_backend_service_iam_binding":       {resourceConverterComputeRegionBackendServiceIamBinding()},
		"google_compute_region_backend_service_iam_member":        {resourceConverterComputeRegionBackendServiceIamMember()},
		"google_compute_disk_iam_policy":                          {resourceConverterComputeDiskIamPolicy()},
		"google_compute_disk_iam_binding":                         {resourceConverterComputeDiskIamBinding()},
		"google_compute_disk_iam_member":                          {resourceConverterComputeDiskIamMember()},
		"google_compute_image_iam_policy":                         {resourceConverterComputeImageIamPolicy()},
		"google_compute_image_iam_binding":                        {resourceConverterComputeImageIamBinding()},
		"google_compute_image_iam_member":                         {resourceConverterComputeImageIamMember()},
		"google_compute_instance_iam_policy":                      {resourceConverterComputeInstanceIamPolicy()},
		"google_compute_instance_iam_binding":                     {resourceConverterComputeInstanceIamBinding()},
		"google_compute_instance_iam_member":                      {resourceConverterComputeInstanceIamMember()},
		"google_compute_region_disk_iam_policy":                   {resourceConverterComputeRegionDiskIamPolicy()},
		"google_compute_region_disk_iam_binding":                  {resourceConverterComputeRegionDiskIamBinding()},
		"google_compute_region_disk_iam_member":                   {resourceConverterComputeRegionDiskIamMember()},
		"google_compute_snapshot_iam_policy":                      {resourceConverterComputeSnapshotIamPolicy()},
		"google_compute_snapshot_iam_binding":                     {resourceConverterComputeSnapshotIamBinding()},
		"google_compute_snapshot_iam_member":                      {resourceConverterComputeSnapshotIamMember()},
		"google_compute_subnetwork_iam_policy":                    {resourceConverterComputeSubnetworkIamPolicy()},
		"google_compute_subnetwork_iam_binding":                   {resourceConverterComputeSubnetworkIamBinding()},
		"google_compute_subnetwork_iam_member":                    {resourceConverterComputeSubnetworkIamMember()},
		"google_data_catalog_entry_group_iam_policy":              {resourceConverterDataCatalogEntryGroupIamPolicy()},
		"google_data_catalog_entry_group_iam_binding":             {resourceConverterDataCatalogEntryGroupIamBinding()},
		"google_data_catalog_entry_group_iam_member":              {resourceConverterDataCatalogEntryGroupIamMember()},
		"google_data_catalog_tag_template_iam_policy":             {resourceConverterDataCatalogTagTemplateIamPolicy()},
		"google_data_catalog_tag_template_iam_binding":            {resourceConverterDataCatalogTagTemplateIamBinding()},
		"google_data_catalog_tag_template_iam_member":             {resourceConverterDataCatalogTagTemplateIamMember()},
		"google_data_fusion_instance_iam_policy":                  {resourceConverterDataFusionInstanceIamPolicy()},
		"google_data_fusion_instance_iam_binding":                 {resourceConverterDataFusionInstanceIamBinding()},
		"google_data_fusion_instance_iam_member":                  {resourceConverterDataFusionInstanceIamMember()},
		"google_dataproc_autoscaling_policy_iam_policy":           {resourceConverterDataprocAutoscalingPolicyIamPolicy()},
		"google_dataproc_autoscaling_policy_iam_binding":          {resourceConverterDataprocAutoscalingPolicyIamBinding()},
		"google_dataproc_autoscaling_policy_iam_member":           {resourceConverterDataprocAutoscalingPolicyIamMember()},
		"google_dataproc_metastore_service_iam_policy":            {resourceConverterDataprocMetastoreServiceIamPolicy()},
		"google_dataproc_metastore_service_iam_binding":           {resourceConverterDataprocMetastoreServiceIamBinding()},
		"google_dataproc_metastore_service_iam_member":            {resourceConverterDataprocMetastoreServiceIamMember()},
		"google_dns_managed_zone_iam_policy":                      {resourceConverterDNSManagedZoneIamPolicy()},
		"google_dns_managed_zone_iam_binding":                     {resourceConverterDNSManagedZoneIamBinding()},
		"google_dns_managed_zone_iam_member":                      {resourceConverterDNSManagedZoneIamMember()},
		"google_gke_hub_membership_iam_policy":                    {resourceConverterGKEHubMembershipIamPolicy()},
		"google_gke_hub_membership_iam_binding":                   {resourceConverterGKEHubMembershipIamBinding()},
		"google_gke_hub_membership_iam_member":                    {resourceConverterGKEHubMembershipIamMember()},
		"google_healthcare_consent_store_iam_policy":              {resourceConverterHealthcareConsentStoreIamPolicy()},
		"google_healthcare_consent_store_iam_binding":             {resourceConverterHealthcareConsentStoreIamBinding()},
		"google_healthcare_consent_store_iam_member":              {resourceConverterHealthcareConsentStoreIamMember()},
		"google_iap_web_iam_policy":                               {resourceConverterIapWebIamPolicy()},
		"google_iap_web_iam_binding":                              {resourceConverterIapWebIamBinding()},
		"google_iap_web_iam_member":                               {resourceConverterIapWebIamMember()},
		"google_iap_tunnel_instance_iam_policy":                   {resourceConverterIapTunnelInstanceIamPolicy()},
		"google_iap_tunnel_instance_iam_binding":                  {resourceConverterIapTunnelInstanceIamBinding()},
		"google_iap_tunnel_instance_iam_member":                   {resourceConverterIapTunnelInstanceIamMember()},
		"google_iap_tunnel_iam_policy":                            {resourceConverterIapTunnelIamPolicy()},
		"google_iap_tunnel_iam_binding":                           {resourceConverterIapTunnelIamBinding()},
		"google_iap_tunnel_iam_member":                            {resourceConverterIapTunnelIamMember()},
		"google_notebooks_instance_iam_policy":                    {resourceConverterNotebooksInstanceIamPolicy()},
		"google_notebooks_instance_iam_binding":                   {resourceConverterNotebooksInstanceIamBinding()},
		"google_notebooks_instance_iam_member":                    {resourceConverterNotebooksInstanceIamMember()},
		"google_notebooks_runtime_iam_policy":                     {resourceConverterNotebooksRuntimeIamPolicy()},
		"google_notebooks_runtime_iam_binding":                    {resourceConverterNotebooksRuntimeIamBinding()},
		"google_notebooks_runtime_iam_member":                     {resourceConverterNotebooksRuntimeIamMember()},
		"google_privateca_ca_pool_iam_policy":                     {resourceConverterPrivatecaCaPoolIamPolicy()},
		"google_privateca_ca_pool_iam_binding":                    {resourceConverterPrivatecaCaPoolIamBinding()},
		"google_privateca_ca_pool_iam_member":                     {resourceConverterPrivatecaCaPoolIamMember()},
		"google_privateca_certificate_template_iam_policy":        {resourceConverterPrivatecaCertificateTemplateIamPolicy()},
		"google_privateca_certificate_template_iam_binding":       {resourceConverterPrivatecaCertificateTemplateIamBinding()},
		"google_privateca_certificate_template_iam_member":        {resourceConverterPrivatecaCertificateTemplateIamMember()},
		"google_pubsub_topic_iam_policy":                          {resourceConverterPubsubTopicIamPolicy()},
		"google_pubsub_topic_iam_binding":                         {resourceConverterPubsubTopicIamBinding()},
		"google_pubsub_topic_iam_member":                          {resourceConverterPubsubTopicIamMember()},
		"google_secret_manager_secret_iam_policy":                 {resourceConverterSecretManagerSecretIamPolicy()},
		"google_secret_manager_secret_iam_binding":                {resourceConverterSecretManagerSecretIamBinding()},
		"google_secret_manager_secret_iam_member":                 {resourceConverterSecretManagerSecretIamMember()},
		"google_scc_source_iam_policy":                            {resourceConverterSecurityCenterSourceIamPolicy()},
		"google_scc_source_iam_binding":                           {resourceConverterSecurityCenterSourceIamBinding()},
		"google_scc_source_iam_member":                            {resourceConverterSecurityCenterSourceIamMember()},
		"google_endpoints_service_iam_policy":                     {resourceConverterServiceManagementServiceIamPolicy()},
		"google_endpoints_service_iam_binding":                    {resourceConverterServiceManagementServiceIamBinding()},
		"google_endpoints_service_iam_member":                     {resourceConverterServiceManagementServiceIamMember()},
		"google_endpoints_service_consumers_iam_policy":           {resourceConverterServiceManagementServiceConsumersIamPolicy()},
		"google_endpoints_service_consumers_iam_binding":          {resourceConverterServiceManagementServiceConsumersIamBinding()},
		"google_endpoints_service_consumers_iam_member":           {resourceConverterServiceManagementServiceConsumersIamMember()},
		"google_vertex_ai_featurestore_iam_policy":                {resourceConverterVertexAIFeaturestoreIamPolicy()},
		"google_vertex_ai_featurestore_iam_binding":               {resourceConverterVertexAIFeaturestoreIamBinding()},
		"google_vertex_ai_featurestore_iam_member":                {resourceConverterVertexAIFeaturestoreIamMember()},
		"google_vertex_ai_featurestore_entitytype_iam_policy":     {resourceConverterVertexAIFeaturestoreEntitytypeIamPolicy()},
		"google_vertex_ai_featurestore_entitytype_iam_binding":    {resourceConverterVertexAIFeaturestoreEntitytypeIamBinding()},
		"google_vertex_ai_featurestore_entitytype_iam_member":     {resourceConverterVertexAIFeaturestoreEntitytypeIamMember()},
		"google_project": {
			resourceConverterProject(),
			resourceConverterProjectBillingInfo(),
		},
		"google_bigtable_instance": {
			resourceConverterBigtableInstance(),
			resourceConverterBigtableCluster(),
		},
		"google_organization_iam_policy":      {resourceConverterOrganizationIamPolicy()},
		"google_organization_iam_binding":     {resourceConverterOrganizationIamBinding()},
		"google_organization_iam_member":      {resourceConverterOrganizationIamMember()},
		"google_organization_policy":          {resourceConverterOrganizationPolicy()},
		"google_project_organization_policy":  {resourceConverterProjectOrgPolicy()},
		"google_folder_iam_policy":            {resourceConverterFolderIamPolicy()},
		"google_folder_iam_binding":           {resourceConverterFolderIamBinding()},
		"google_folder_iam_member":            {resourceConverterFolderIamMember()},
		"google_folder_organization_policy":   {resourceConverterFolderOrgPolicy()},
		"google_kms_crypto_key_iam_policy":    {resourceConverterKmsCryptoKeyIamPolicy()},
		"google_kms_crypto_key_iam_binding":   {resourceConverterKmsCryptoKeyIamBinding()},
		"google_kms_crypto_key_iam_member":    {resourceConverterKmsCryptoKeyIamMember()},
		"google_kms_key_ring_iam_policy":      {resourceConverterKmsKeyRingIamPolicy()},
		"google_kms_key_ring_iam_binding":     {resourceConverterKmsKeyRingIamBinding()},
		"google_kms_key_ring_iam_member":      {resourceConverterKmsKeyRingIamMember()},
		"google_project_iam_policy":           {resourceConverterProjectIamPolicy()},
		"google_project_iam_binding":          {resourceConverterProjectIamBinding()},
		"google_project_iam_member":           {resourceConverterProjectIamMember()},
		"google_project_iam_custom_role":      {resourceConverterProjectIAMCustomRole()},
		"google_organization_iam_custom_role": {resourceConverterOrganizationIAMCustomRole()},
		"google_vpc_access_connector":         {resourceConverterVPCAccessConnector()},
	}
}