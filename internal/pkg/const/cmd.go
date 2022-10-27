package consts

const (
	CreateImage = `qemu-img create -f qcow2 \
		-o cluster_size=2M,backing_file=%s \
		%s 60G
	`

	CreateVm = `virt-install \
		--connect qemu:///system \
		--virt-type kvm \
		--name %s \
		--vcpus %d \
		--ram %d \
		--disk path=/home/aaron/zagent/kvm/image/test-1.qcow2,size=%d,sparse \
		--network default \
		--boot hd \
		--os-variant generic \
		--noreboot \
		--wait -1 \
		--force 
	`

	StartVm = `virsh --connect qemu:///system start %s`
)

var ()
