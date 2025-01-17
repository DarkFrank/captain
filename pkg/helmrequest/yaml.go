package helmrequest

var helmRequestCRDYaml = `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: helmrequests.app.alauda.io
spec:
  group: app.alauda.io
  conversion:
    strategy: None
  names:
    kind: HelmRequest
    listKind: HelmRequestList
    plural: helmrequests
    singular: helmrequest
    shortNames:
      - hr
      - hrs
  scope: Namespaced
  versions:
  - name: v1alpha1
    additionalPrinterColumns:
    - name: Chart
      type: string
      description: The chart of this HelmRequest
      jsonPath: .spec.chart
    - name: Version
      type: string
      description: Version of this chart
      jsonPath: .spec.version
    - name: Namespace
      type: string
      description: The namespace which the chart deployed to
      jsonPath: .spec.namespace
    - name: AllCluster
      type: boolean
      description: Is this chart will be installed to all cluster
      jsonPath: .spec.installToAllClusters
    - name: Phase
      type: string
      description: The phase of this HelmRequest
      jsonPath: .status.phase
    - name: Age
      type: date
      jsonPath: .metadata.creationTimestamp
    schema:
      openAPIV3Schema:
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            properties:
              chart:
                type: string
              clusterName:
                description: ClusterName is the cluster where the chart will be installed.
                  If InstallToAllClusters=true, this field will be ignored
                type: string
              dependencies:
                description: Dependencies is the dependencies of this HelmRequest,
                  it's a list of there names THe dependencies must lives in the same
                  namespace, and each of them must be in Synced status before we sync
                  this HelmRequest
                items:
                  type: string
                type: array
              installToAllClusters:
                description: InstallToAllClusters will install this chart to all available
                  clusters, even the cluster was created after this chart. If this
                  field is true, ClusterName will be ignored(useless)
                type: boolean
              namespace:
                description: Namespace is the namespace where the Release object will
                  be lived in. Notes this should be used with the values defined in
                  the chart， otherwise the install will failed
                type: string
              releaseName:
                description: ReleaseName is the Release name to be generated, default
                  to HelmRequest.Name. If we want to manually install this chart to
                  multi clusters, we may have different HelmRequest name(with cluster
                  prefix or suffix) and same release name
                type: string
              values:
                description: Values represents a collection of chart values.
                type: object
                nullable: true
                x-kubernetes-preserve-unknown-fields: true
              valuesFrom:
                description: ValuesFrom represents values from ConfigMap/Secret...
                items:
                  description: ValuesFromSource represents a source of values, only
                    one of it's fields may be set
                  properties:
                    configMapKeyRef:
                      description: ConfigMapKeyRef selects a key of a ConfigMap
                      properties:
                        key:
                          description: The key to select.
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            TODO: Add other useful fields. apiVersion, kind, uid?'
                          type: string
                        optional:
                          description: Specify whether the ConfigMap or its key must
                            be defined
                          type: boolean
                      required:
                      - key
                      type: object
                    secretKeyRef:
                      description: SecretKeyRef selects a key of a Secret
                      properties:
                        key:
                          description: The key of the secret to select from.  Must
                            be a valid secret key.
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            TODO: Add other useful fields. apiVersion, kind, uid?'
                          type: string
                        optional:
                          description: Specify whether the Secret or its key must
                            be defined
                          type: boolean
                      required:
                      - key
                      type: object
                  type: object
                type: array
              version:
                type: string
            type: object
          status:
            properties:
              conditions:
                items:
                  properties:
                    lastProbeTime:
                      description: Last time we probed the condition.
                      format: date-time
                      type: string
                      nullable: true
                    lastTransitionTime:
                      description: Last time the condition transitioned from one status
                        to another.
                      format: date-time
                      type: string
                      nullable: true
                    message:
                      description: Human-readable message indicating details about
                        last transition.
                      type: string
                    reason:
                      description: Unique, one-word, CamelCase reason for the condition's
                        last transition.
                      type: string
                    status:
                      description: 'Status is the status of the condition. Can be
                        True, False, Unknown. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions'
                      type: string
                    type:
                      description: 'Type is the type of the condition. More info:
                        https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions'
                      type: string
                  type: object
                type: array
              lastSpecHash:
                description: LastSpecHash store the has value of the synced spec,
                  if this value not equal to the current one, means we need to do
                  a update for the chart
                type: string
              notes:
                description: Notes is the contents from helm (after helm install successfully
                  it will be printed to the console
                type: string
              phase:
                description: HelmRequestPhase is a label for the condition of a HelmRequest
                  at the current time.
                type: string
              reason:
                description: Reason will store the reason why the HelmRequest deploy
                  failed
                type: string
              syncedClusters:
                description: SyncedClusters will store the synced clusters if InstallToAllClusters
                  is true
                items:
                  type: string
                type: array
              version:
                description: Verions is the real version that installed
                type: string
            type: object
        required:
        - spec
        type: object
    served: true
    storage: true
    subresources:
      status: {}
`
