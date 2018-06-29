// +build !linux

package daemon

func clobberDefaultAppArmorProfile() error {
	return nil
}

func ensureDefaultAppArmorProfile() error {
	return nil
}
