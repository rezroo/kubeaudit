package cmd

// KubeauditConfig sets up config for kubeaudit from flag `config`
type KubeauditConfig struct {
	APIVersion string               `yaml:"apiVersion"`
	Kind       string               `yaml:"kind"`
	Spec       *KubeauditConfigSpec `yaml:"spec"`
	Audit      bool                 `yaml:"audit"`
}

// KubeauditConfigSpec contains Config Spec
type KubeauditConfigSpec struct {
	Manifest     []*KubeauditConfigManifest   `yaml:"manifest"`
	Capabilities *KubeauditConfigCapabilities `yaml:"capabilities"`
	Overrides    *KubeauditConfigOverrides    `yaml:"overrides"`
}

// KubeauditConfigManifest contains path to the manifests to audit
type KubeauditConfigManifest struct {
	Path string `yaml:"path"`
}

// KubeauditConfigCapabilities contains list of capabilities supported
type KubeauditConfigCapabilities struct {
	NetAdmin       string `yaml:"NET_ADMIN"`
	SetPCAP        string `yaml:"SETPCAP"`
	MKNOD          string `yaml:"MKNOD"`
	AuditWrite     string `yaml:"AUDIT_WRITE"`
	Chown          string `yaml:"CHOWN"`
	NetRaw         string `yaml:"NET_RAW"`
	DacOverride    string `yaml:"DAC_OVERRIDE"`
	FOWNER         string `yaml:"FOWNER"`
	FSetID         string `yaml:"FSETID"`
	Kill           string `yaml:"KILL"`
	SetGID         string `yaml:"SETGID"`
	SetUID         string `yaml:"SETUID"`
	NetBindService string `yaml:"NET_BIND_SERVICE"`
	SYSChroot      string `yaml:"SYS_CHROOT"`
	SetFCAP        string `yaml:"SETFCAP"`
}

// KubeauditConfigOverrides contains list of available overrides
type KubeauditConfigOverrides struct {
	PrivilegeEscalation                string `yaml:"privilege-escalation"`
	Privileged                         string `yaml:"privileged"`
	RunAsRoot                          string `yaml:"run-as-root"`
	AutomountServiceAccountToken       string `yaml:"automount-service-account-token"`
	ReadOnlyRootFilesystemFalse        string `yaml:"read-only-root-filesystem-false"`
	NonDefaultDenyIngressNetworkPolicy string `yaml:"non-default-deny-ingress-network-policy"`
	NonDefaultDenyEgressNetworkPolicy  string `yaml:"non-default-deny-egress-network-policy"`
}

func mapOverridesToStructFields(label string) string {
	if label == "allow-privilege-escalation" {
		return "PrivilegeEscalation"
	}
	if label == "allow-privileged" {
		return "Privileged"
	}
	if label == "allow-run-as-root" {
		return "RunAsRoot"
	}
	if label == "allow-automount-service-account-token" {
		return "AutomountServiceAccountToken"
	}
	if label == "allow-read-only-root-filesystem-false" {
		return "ReadOnlyRootFilesystemFalse"
	}
	if label == "allow-non-default-deny-egress-network-policy" {
		return "NonDefaultDenyEgressNetworkPolicy"
	}
	if label == "allow-non-default-deny-ingress-network-policy" {
		return "NonDefaultDenyIngressNetworkPolicy"
	}
	return ""
}
