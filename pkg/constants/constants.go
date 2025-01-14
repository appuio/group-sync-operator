package constants

const (
	AnnotationBase        = "group-sync-operator.redhat-cop.io"
	SyncTimestamp         = AnnotationBase + "/sync-time"
	SyncSourceURL         = AnnotationBase + "/sync.source.url"
	SyncSourceHost        = AnnotationBase + "/sync.source.host"
	SyncSourceUID         = AnnotationBase + "/sync.source.uid"
	SyncProvider          = AnnotationBase + "/sync-provider"
	SyncProviderNamespace = AnnotationBase + "/sync-provider.namespace"
	HierarchyChildren     = "hierarchy_children"
	HierarchyParent       = "hierarchy_parent"
	HierarchyParents      = "hierarchy_parents"
)
