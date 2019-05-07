package drm

type Device struct {
	fd uintptr
}

func NewDevice(fd uintptr) *Device {
	return &Device{fd}
}

type VersionResp struct {
	Major, Minor, Patch int32
	Name, Date, Desc string
}

func (d *Device) Version() (*VersionResp, error) {
	var v versionResp
	if err := version(d.fd, &v); err != nil {
		return nil, err
	}

	name := allocBytes(&v.name, v.nameLen)
	date := allocBytes(&v.date, v.dateLen)
	desc := allocBytes(&v.desc, v.descLen)

	if err := version(d.fd, &v); err != nil {
		return nil, err
	}

	return &VersionResp{
		Major: v.major,
		Minor: v.minor,
		Patch: v.patch,
		Name: string(name),
		Date: string(date),
		Desc: string(desc),
	}, nil
}
