package blockchain

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"time"

	"github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/prysm/beacon-chain/db/filters"
	"github.com/prysmaticlabs/prysm/shared/bytesutil"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/sirupsen/logrus"
	"github.com/awalterschulze/gographviz"
)

const latestSlotCount = 10

type node struct {
	parentRoot string
	selfRoot string
}

// BlockTreeHandler is a handler to serve /tree page in metrics.
func (s *Service) BlockTreeHandler(w http.ResponseWriter, _ *http.Request) {
	buf := new(bytes.Buffer)

	if _, err := fmt.Fprintf(w, "\n %s\t%s\t", "Head slot", "Head root"); err != nil {
		logrus.WithError(err).Error("Failed to render chain heads page")
		return
	}

	if _, err := fmt.Fprintf(w, "\n %s\t%s\t", "---------", "---------"); err != nil {
		logrus.WithError(err).Error("Failed to render chain heads page")
		return
	}

	slots := s.latestHeadSlots()
	for _, slot := range slots {
		r := hex.EncodeToString(bytesutil.Trunc(s.canonicalRoots[uint64(slot)]))
		if _, err := fmt.Fprintf(w, "\n %d\t\t%s\t", slot, r); err != nil {
			logrus.WithError(err).Error("Failed to render chain heads page")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(buf.Bytes()); err != nil {
		log.WithError(err).Error("Failed to render chain heads page")
	}
}

// HeadsHandler is a handler to serve /heads page in metrics.
func (s *Service) HeadsHandler(w http.ResponseWriter, r *http.Request) {
	graphAst, _ := gographviz.Parse([]byte(`digraph G{}`))
	graph := gographviz.NewGraph()
	gographviz.Analyse(graphAst, graph)

	currentSlot := s.currentSlot()
	filter := filters.NewFilter().SetStartSlot(1).SetEndSlot(currentSlot)
	blks, err := s.beaconDB.Blocks(context.Background(), filter)
	nodes := make([]node, len(blks))
	for i:=0; i<len(nodes) ; i++  {
		r, _ := ssz.SigningRoot(blks[i])
		nodes[i].selfRoot = hex.EncodeToString(r[:1])
		nodes[i].parentRoot = hex.EncodeToString(blks[i].ParentRoot[:1])
	}
	for i := 0; i<len(nodes); i++ {
		n := nodes[i]
		if err := graph.AddNode("G", n.selfRoot, nil); err != nil {
			log.WithError(err).Error("Could not add node to graph")
		}
		if i != 0 {
			if err := graph.AddEdge(n.parentRoot, n.selfRoot, true, nil); err != nil {
				log.WithError(err).Error("Could not add node to graph")
			}
		}
	}

	img, err := dotToImage([]byte(graph.String()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, img)
}

// This returns the latest head slots in a slice and up to latestSlotCount
func (s *Service) latestHeadSlots() []int {
	slots := make([]int, 0, len(s.canonicalRoots))
	for k := range s.canonicalRoots {
		slots = append(slots, int(k))
	}
	sort.Ints(slots)
	if (len(slots)) > latestSlotCount {
		return slots[len(slots)-latestSlotCount:]
	}
	return slots
}

func dotToImage(dot []byte) (string, error) {
	format := "svg"
	dotExe, err := exec.LookPath("dot")
	if err != nil {
		log.Fatalln("unable to find program 'dot', please install it or check your PATH")
	}

	img := filepath.Join(os.TempDir(), fmt.Sprintf("go-callvis_export.%s", format))

	cmd := exec.Command(dotExe, fmt.Sprintf("-T%s", format), "-o", img)
	cmd.Stdin = bytes.NewReader(dot)
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return img, nil
}

func (s *Service) currentSlot() uint64 {
	diff := time.Now().Unix() - s.GenesisTime().Unix()
	return uint64(diff) / params.BeaconConfig().SecondsPerSlot
}
