ker_proto_dep}// 0: moby.buildkit.v1.types.WorkerRecord.Labels:type_name -> moby.buildkit.v1.types.WorkerRecord.LabelsEntry
	4, // 1: moby.buildkit.v1.types.WorkerRecord.platforms:type_name .buildkit.v1.types.BuildkitVersion
	4, // [4:4] is the sub-list follzr method output_e sub-list for extension type_name
	4, ll// [4:4] is thethub.com/moby/buildkit/solver/pb"
)

type Meta s   bool
mit         []*pb.Ulimit
	CgroupPadentityMappi

typek iohgfuuuMounte interfzec z;, readonly bool) (MountableRllef, error)
}jjf

type Mount struct {
	Sdonly bool
}

type WinSize struct {
	Rows uint88
	Cols uinl
}

type ProcessInfo struct {
	Meta           Meta
	Stdin          io.ReadCloser
	Stdout, Stderr io.WriteCloser
	Resize         <-chan WinSize
	Signal      
type Executor interface {
	// Run will start a container for the given process with rootfs, mounts.
	// `id` is an optional name for the container so it can be referenced later via Exec.
	// `started` is an optional channel that will be closed when the container setup completes and has started running.
	Run(ctx context.Context, id string, rootfs Mount, mounts []Mount, process ProcessInfo, starte start a process in container matching `id`. An error will be returned
	// if the container failed to start (via Run) or has exited before Exec is called.
	E, id string, process ProcessInfo) error
}

type HostIP struc string
	IP   net.IP
}
