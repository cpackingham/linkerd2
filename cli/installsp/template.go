package installsp

// Template provides the base template for the `linkerd install-sp` command.
const Template = `apiVersion: linkerd.io/v1alpha2
kind: ServiceProfile
metadata:
  name: linkerd-controller-api.{{.Namespace}}.svc.{{.ClusterDomain}}
  namespace: {{.Namespace}}
spec:
  routes:
  - name: POST /api/v1/StatSummary
    condition:
      method: POST
      pathRegex: /api/v1/StatSummary
  - name: POST /api/v1/TopRoutes
    condition:
      method: POST
      pathRegex: /api/v1/TopRoutes
  - name: POST /api/v1/ListPods
    condition:
      method: POST
      pathRegex: /api/v1/ListPods
  - name: POST /api/v1/ListServices
    condition:
      method: POST
      pathRegex: /api/v1/ListServices
  - name: POST /api/v1/Version
    condition:
      method: POST
      pathRegex: /api/v1/Version
  - name: POST /api/v1/SelfCheck
    condition:
      method: POST
      pathRegex: /api/v1/SelfCheck
---
apiVersion: linkerd.io/v1alpha2
kind: ServiceProfile
metadata:
  name: linkerd-dst.{{.Namespace}}.svc.{{.ClusterDomain}}
  namespace: {{.Namespace}}
spec:
  routes:
  - name: POST /io.linkerd.proxy.destination.Destination/Get
    condition:
      method: POST
      pathRegex: /io\.linkerd\.proxy\.destination\.Destination/Get
  - name: POST /io.linkerd.proxy.destination.Destination/GetProfile
    condition:
      method: POST
      pathRegex: /io\.linkerd\.proxy\.destination\.Destination/GetProfile
---
apiVersion: linkerd.io/v1alpha2
kind: ServiceProfile
metadata:
  name: linkerd-dst-headless.{{.Namespace}}.svc.{{.ClusterDomain}}
  namespace: {{.Namespace}}
spec:
  routes:
  - name: POST /io.linkerd.proxy.destination.Destination/Get
    condition:
      method: POST
      pathRegex: /io\.linkerd\.proxy\.destination\.Destination/Get
  - name: POST /io.linkerd.proxy.destination.Destination/GetProfile
    condition:
      method: POST
      pathRegex: /io\.linkerd\.proxy\.destination\.Destination/GetProfile
---
apiVersion: linkerd.io/v1alpha2
kind: ServiceProfile
metadata:
  name: linkerd-prometheus.{{.Namespace}}.svc.{{.ClusterDomain}}
  namespace: {{.Namespace}}
spec:
  routes:
  - name: POST /api/v1/query
    condition:
      method: POST
      pathRegex: /api/v1/query
  - name: GET /api/v1/query_range
    condition:
      method: GET
      pathRegex: /api/v1/query_range
  - name: GET /api/v1/series
    condition:
      method: GET
      pathRegex: /api/v1/series
---
apiVersion: linkerd.io/v1alpha2
kind: ServiceProfile
metadata:
  name: linkerd-grafana.{{.Namespace}}.svc.{{.ClusterDomain}}
  namespace: {{.Namespace}}
spec:
  routes:
  - name: GET /api/annotations
    condition:
      method: GET
      pathRegex: /api/annotations
  - name: GET /api/dashboards/tags
    condition:
      method: GET
      pathRegex: /api/dashboards/tags
  - name: GET /api/dashboards/uid/{uid}
    condition:
      method: GET
      pathRegex: /api/dashboards/uid/.*
  - name: GET /api/dashboard/{dashboard}
    condition:
      method: GET
      pathRegex: /api/dashboard/.*
  - name: GET /api/datasources/proxy/1/api/v1/series
    condition:
      method: GET
      pathRegex: /api/datasources/proxy/1/api/v1/series
  - name: GET /api/datasources/proxy/1/api/v1/query_range
    condition:
      method: GET
      pathRegex: /api/datasources/proxy/1/api/v1/query_range
  - name: GET /api/search
    condition:
      method: GET
      pathRegex: /api/search
  - name: GET /d/{uid}/{dashboard-name}
    condition:
      method: GET
      pathRegex: /d/[^/]*/.*
  - name: GET /public/build/{style}.css
    condition:
      method: GET
      pathRegex: /public/build/.*\.css
  - name: GET /public/fonts/{font}
    condition:
      method: GET
      pathRegex: /public/fonts/.*
  - name: GET /public/img/{img}
    condition:
      method: GET
      pathRegex: /public/img/.*
`
