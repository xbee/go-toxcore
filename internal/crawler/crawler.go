// implementation tox-bootstrap in golang
package main

import (
	"fmt"
	"gopp"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kitech/colog"
	"github.com/kitech/go-toxcore/internal"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	colog.Register()
}

/* Seconds to wait between new crawler instances */
const NEW_CRAWLER_INTERVAL = 180

/* Maximum number of concurrent crawler instances */
const MAX_CRAWLERS = 6

/* Number of seconds to wait for new nodes before a crawler times out and exits */
const CRAWLER_TIMEOUT = 20

/* Default maximum number of nodes the nodes list can store */
const DEFAULT_NODES_LIST_SIZE = 4096

/* Seconds to wait between getnodes requests */
const GETNODES_REQUEST_INTERVAL = 1

/* Max number of nodes to send getnodes requests to per GETNODES_REQUEST_INTERVAL */
const MAX_GETNODES_REQUESTS = 5

/* Number of random node requests to make for each node we send a request to */
const NUM_RAND_GETNODE_REQUESTS = 32

type Crawler struct {
	// tox                   *tox.Tox
	dht                   *toxin.DHT
	nodes_list            *toxin.NodeFormatList
	num_nodes             uint32
	nodes_list_size       uint32
	send_ptr              uint32    /* index of the oldest node that we haven't sent a getnodes request to */
	last_new_node         time.Time /* Last time we found an unknown node */
	last_getnodes_request time.Time

	// pthread_t      tid;
	// pthread_attr_t attr;
}

const NUM_BOOTSTRAP_NODES = 14
const NUM_BOOTSTRAPS = 4

type toxNodes struct {
	ip   string
	port uint16
	key  string
}

var bs_nodes = []toxNodes{
	{"198.98.51.198", 33445, "1D5A5F2F5D6233058BF0259B09622FB40B482E4FA0931EB8FD3AB8E7BF7DAF6F"},
	{"130.133.110.14", 33445, "461FA3776EF0FA655F1A05477DF1B3B614F7D6B124F7DB1DD4FE3C08B03B640F"},
	{"205.185.116.116", 33445, "A179B09749AC826FF01F37A9613F6B57118AE014D4196A0E1105A98F93A54702"},
	{"51.254.84.212", 33445, "AEC204B9A4501412D5F0BB67D9C81B5DB3EE6ADA64122D32A3E9B093D544327D"},
	{"5.135.59.163", 33445, "2D320F971EF2CA18004416C2AAE7BA52BF7949DB34EA8E2E21AF67BD367BE211"},
	{"185.58.206.164", 33445, "24156472041E5F220D1FA11D9DF32F7AD697D59845701CDD7BE7D1785EB9DB39"},
	{"194.249.212.109", 33445, "3CEE1F054081E7A011234883BC4FC39F661A55B73637A5AC293DDF1251D9432B"},
	{"92.54.84.70", 33445, "5625A62618CB4FCA70E147A71B29695F38CC65FF0CBD68AD46254585BE564802"},
	{"95.215.46.114", 33445, "5823FB947FF24CF83DDFAC3F3BAA18F96EA2018B16CC08429CB97FA502F40C23"},
	{"5.189.176.217", 5190, "2B2137E094F743AC8BD44652C55F41DFACC502F125E99E4FE24D40537489E32F"},
	{"136.243.141.187", 443, "6EE1FADE9F55CC7938234CC07C864081FC606D8FE7B751EDA217F268F1078A39"},
	{"37.187.122.30", 33445, "BEB71F97ED9C99C04B8489BB75579EB4DC6AB6F441B603D63533122F1858B51D"},
	{"85.21.144.224", 33445, "8F738BBC8FA9394670BCAB146C67A507B9907C8E564E28C2B59BEBB2FF68711B"},
	{"144.217.86.39", 33445, "7E5668E0EE09E19F320AD47902419331FFEE147BB3606769CFBE921A2A2FD34C"},
	{"37.97.185.116", 5190, "E59A0E71ADA20D35BD1B0957059D7EF7E7792B3D680AE25C6F4DBBA09114D165"},
	// {NULL, 0, NULL},
}

func bootstrap_dht(cwl *Crawler) {
	dht := cwl.dht
	for i := 0; i < NUM_BOOTSTRAPS; i++ {
		r := rand.Int() % NUM_BOOTSTRAP_NODES
		n := bs_nodes[r]
		rc := dht.BootstrapFromAddress(n.ip, false, n.port, n.key)
		if !rc {
			log.Println(rc, n)
		}
	}
}

var FLAG_EXIT bool = false

func node_crawled(cwl *Crawler, pubkey string) bool {
	for i := 0; i < int(cwl.num_nodes); i++ {
		node := cwl.nodes_list.Get(i)
		if node.Pubkey() == pubkey {
			return true
		}
	}
	return false
}

func cb_getnodes_response(ip string, port uint16, pubkey string, object interface{}) {
	log.Println(fmt.Sprintf("%s:%d", ip, port), len(pubkey), pubkey, object == nil)
	cwl := object.(*Crawler)

	if node_crawled(cwl, pubkey) {
		log.Println("already crawled:", cwl.num_nodes)
		return
	}

	if cwl.num_nodes+1 >= cwl.nodes_list_size {
		n, ok := cwl.nodes_list.Expand()
		if !ok {
			log.Println("Expand error")
			return
		}
		cwl.nodes_list_size = uint32(n)
	}

	node := cwl.nodes_list.Get(int(cwl.num_nodes))
	node.Set(ip, port, pubkey)

	cwl.num_nodes++
	log.Println("num_nodes:", cwl.num_nodes)
	// 难道说是内网中的节点搜索不到？
	wantkeys := []string{
	// "A179B09749AC826FF01F37A9613F6B57118AE014D4196A0E1105A98F93A54702",
	}
	for idx, key := range wantkeys {
		if len(key) != len(pubkey) {
			log.Fatalln("wtf:", idx, key, pubkey)
		}
		if key == pubkey {
			log.Println("oh holy shit", idx, len(key), len(pubkey))
			time.Sleep(5 * time.Second)
			log.Fatalln("oh holy shit", idx)
		}
	}
	if ip == "205.185.116.116" { // test ok, get  right pubkey
		// log.Fatalln(ip, port, pubkey)
	}
}

func send_node_requests(cwl *Crawler) int {

	count := 0
	var i int

	gopp.KeepMe()
	dht := cwl.dht

	for i = int(cwl.send_ptr); count < MAX_GETNODES_REQUESTS && i < int(cwl.num_nodes); i++ {
		nf := cwl.nodes_list.Get(i)
		dht.GetNodes(nf.IPPortC(), nf.PubkeyC(), nf.PubkeyC())

		for j := 0; j < NUM_RAND_GETNODE_REQUESTS; j++ {
			r := rand.Uint32() % cwl.num_nodes

			nf2 := cwl.nodes_list.Get(int(r))
			dht.GetNodes(nf.IPPortC(), nf.PubkeyC(), nf2.PubkeyC())
		}

		count++
	}

	if count == 0 {
		log.Println("wtf", "send ptr:", i, "count:", count)
	}
	cwl.send_ptr = uint32(i)
	cwl.last_getnodes_request = get_time()

	return count
}

func NewCrawler(dht *toxin.DHT) *Crawler {
	cwl := &Crawler{}

	cwl.dht = dht
	cwl.nodes_list = toxin.NewNodeFormatN(DEFAULT_NODES_LIST_SIZE)
	cwl.nodes_list_size = DEFAULT_NODES_LIST_SIZE

	cwl.dht.CallbackGetnodesResponse(cb_getnodes_response, cwl)

	cwl.last_getnodes_request = get_time()
	cwl.last_new_node = get_time()

	// bootstrap_dht(cwl)
	return cwl
}

/// util
func get_time() time.Time { return time.Now() }
func timed_out(ts time.Time, timeout int) bool {
	return false
}

///
func main() {
	var ip toxin.IP
	(&ip).Init()

	netcore := toxin.NewNetworkCore(ip, 12345)
	dht := toxin.NewDHT(netcore)

	onion := toxin.NewOnion(dht)
	onion_an := toxin.NewOnionAnnounce(dht)
	if onion == nil || onion_an == nil {
		log.Println("error: ", "Couldn't initialize Tox Onion. Exiting.")
		os.Exit(-1)
	}
	log.Println(onion, onion_an)

	secret_key := dht.SecretKey()
	tcpsrv := toxin.NewTCPServer(true, []uint16{33445}, secret_key, onion)
	log.Println(tcpsrv)

	dht.Dump()

	cwl := NewCrawler(dht)
	bootstrap_dht(cwl)

	dht.LANdiscoveryInit()

	go func() {
		for {
			send_node_requests(cwl)
			dht.Do()
			tcpsrv.Do()
			netcore.Poll()
			if toxin.NeedDump(dht) {
				dht.Dump()
				log.Println()
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	select {}

}
