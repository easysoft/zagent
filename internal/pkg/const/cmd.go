package consts

const (
	CmdCreateImage = `qemu-img create -f qcow2 \
-o cluster_size=2M,backing_file=%s \
%s 60G
`

	CmdCreateVm = `virt-install \
--connect qemu:///system \
--virt-type kvm \
--name %s \
--vcpus %d \
--ram %d \
--disk path=%s,size=%d,sparse \
--network default \
--boot hd \
--os-variant generic \
--noreboot \
--wait -1 \
--force 
`

	CmdStartVm = `virsh --connect qemu:///system start %s`

	CmdExportVm = `qemu-img convert -f qcow2 -O qcow2 %s.qcow2 %s.qcow2`
)

var ()
