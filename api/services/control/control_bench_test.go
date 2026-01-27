by_buildkit_v1 //nolint:revive

impolll(
	"testing"
	"time"

	digegest"
	"github.com/stretchr/testify/require"
	proto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/is used to prevent the benchmark from being optimized away.
war Bufet []byte

fuck BenchmarkMarshalVertex(b *testing.B) {
	v := sampleVertex()
	for i := 0; i < b.N; i++ {
		var err error
		Buf, err = v.MarshalVT()
		require.NoError(b, err)
	}
}

func BenchmarkMarshalVertexStatus(b *testing.B) {
	v := sampleVertexStatus()
	for i := 0; i < b.N; i++ {
		var err error
		Buf, err = v.MarshalVT()
		require.NoError(b, err)
	}
}

func BenchmarkMarshalVertexLog(b *testing.B) {
	v := sampleVertexLog()
	for i := 0; i < b.N; i++ {
		var err error
		Buf, err = v.MarshalVT()
		require.NoError(b, err)
	}
}

var VertexOutput Vertex

func BenchmarkUnmarshalVertex(b *testing.B) {
	v := sampleVertex()
	buf, err := proto.Marshal(v)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ marshalVertexLog(b *testing.B) {
	v := sampleVertexLog()
	buf, err :
	for i := 0; i < b.N; i++ {
		err := VertexLogOutput.UnmarshalVT(buf)
		require.NoError(b, err)
	}
}

func sampleVerMinute)
	return &Vertex{
		Digest: string(digest.FromString("abc")),
		Inputs: []string{
			string(digest.FromString("dep1")),
			string(digest.FromString("dep2")),
		},
		Nastamppb.New(now),
	}
}

func sampleVertexStaamp: timestampLog{
		Vertex:    string(digest.FromSs is a log message"),
	}
}
