# trivy-renderer

If you use the [Trivy Operator](https://github.com/aquasecurity/trivy-operator) in your cluster, it will save all
the reports for the scans as CRDs like `vulnerabilityreports.aquasecurity.github.io`.

Scrolling through these big YAML definitions is not fun, and if you don't have or don't need a fancy GUI,
you can use this tool to get a basic ASCII table of your reports.

Just run the tool, and it will load your kubernetes config (either $KUBECONFIG or ~/.kube/config)
and display all vulnerabilityreports for the currently selected namespace.

## Example

```
% ./trivy-renderer                                                                                                                                                                                                                                                                clancy/trivy-system
Name: replicaset-7449bfff9f, Namespace: cert-manager
┌─────────────────────┬────────────────┬──────────┬───────────┬────────────────┬────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│      RESOURCE       │      CVE       │ SEVERITY │ INSTALLED │     FIXED      │                                                        META                                                        │
├─────────────────────┼────────────────┼──────────┼───────────┼────────────────┼────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ golang.org/x/crypto │ CVE-2025-22869 │ HIGH     │ v0.31.0   │ 0.35.0         │ golang.org/x/crypto/ssh: Denial of Service in the Key Exchange of golang.org/x/crypto/ssh                          │
│ golang.org/x/net    │ CVE-2025-22870 │ MEDIUM   │ v0.33.0   │ 0.36.0         │ golang.org/x/net/proxy: golang.org/x/net/http/httpproxy: HTTP Proxy bypass using IPv6 Zone IDs in golang.org/x/net │
│ golang.org/x/net    │ CVE-2025-22872 │ MEDIUM   │ v0.33.0   │ 0.38.0         │ golang.org/x/net/html: Incorrect Neutralization of Input During Web Page Generation in x/net in golang.org/x/net   │
│ stdlib              │ CVE-2025-22871 │ MEDIUM   │ v1.23.6   │ 1.23.8, 1.24.2 │ net/http: Request smuggling due to acceptance of invalid chunked data in net/http                                  │
└─────────────────────┴────────────────┴──────────┴───────────┴────────────────┴────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Name: replicaset-cert-manager-584674cf4-cert-manager-controller, Namespace: cert-manager
┌───────────────────────────────┬────────────────┬──────────┬───────────┬────────────────┬────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│           RESOURCE            │      CVE       │ SEVERITY │ INSTALLED │     FIXED      │                                                        META                                                        │
├───────────────────────────────┼────────────────┼──────────┼───────────┼────────────────┼────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ github.com/go-jose/go-jose/v4 │ CVE-2025-27144 │ MEDIUM   │ v4.0.2    │ 4.0.5          │ go-jose: Go JOSE's Parsing Vulnerable to Denial of Service                                                         │
│ github.com/golang-jwt/jwt/v5  │ CVE-2025-30204 │ HIGH     │ v5.2.1    │ 5.2.2          │ golang-jwt/jwt: jwt-go allows excessive memory allocation during header parsing                                    │
│ golang.org/x/crypto           │ CVE-2025-22869 │ HIGH     │ v0.31.0   │ 0.35.0         │ golang.org/x/crypto/ssh: Denial of Service in the Key Exchange of golang.org/x/crypto/ssh                          │
│ golang.org/x/net              │ CVE-2025-22870 │ MEDIUM   │ v0.33.0   │ 0.36.0         │ golang.org/x/net/proxy: golang.org/x/net/http/httpproxy: HTTP Proxy bypass using IPv6 Zone IDs in golang.org/x/net │
│ golang.org/x/net              │ CVE-2025-22872 │ MEDIUM   │ v0.33.0   │ 0.38.0         │ golang.org/x/net/html: Incorrect Neutralization of Input During Web Page Generation in x/net in golang.org/x/net   │
│ stdlib                        │ CVE-2025-22871 │ MEDIUM   │ v1.23.6   │ 1.23.8, 1.24.2 │ net/http: Request smuggling due to acceptance of invalid chunked data in net/http                                  │
└───────────────────────────────┴────────────────┴──────────┴───────────┴────────────────┴────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
Name: replicaset-cert-manager-webhook-68f89ddf6c-cert-manager-webhook, Namespace: cert-manager
┌─────────────────────┬────────────────┬──────────┬───────────┬────────────────┬────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│      RESOURCE       │      CVE       │ SEVERITY │ INSTALLED │     FIXED      │                                                        META                                                        │
├─────────────────────┼────────────────┼──────────┼───────────┼────────────────┼────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┤
│ golang.org/x/crypto │ CVE-2025-22869 │ HIGH     │ v0.31.0   │ 0.35.0         │ golang.org/x/crypto/ssh: Denial of Service in the Key Exchange of golang.org/x/crypto/ssh                          │
│ golang.org/x/net    │ CVE-2025-22870 │ MEDIUM   │ v0.33.0   │ 0.36.0         │ golang.org/x/net/proxy: golang.org/x/net/http/httpproxy: HTTP Proxy bypass using IPv6 Zone IDs in golang.org/x/net │
│ golang.org/x/net    │ CVE-2025-22872 │ MEDIUM   │ v0.33.0   │ 0.38.0         │ golang.org/x/net/html: Incorrect Neutralization of Input During Web Page Generation in x/net in golang.org/x/net   │
│ stdlib              │ CVE-2025-22871 │ MEDIUM   │ v1.23.6   │ 1.23.8, 1.24.2 │ net/http: Request smuggling due to acceptance of invalid chunked data in net/http                                  │
└─────────────────────┴────────────────┴──────────┴───────────┴────────────────┴────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```
