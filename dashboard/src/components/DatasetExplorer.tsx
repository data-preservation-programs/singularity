import React, { useEffect, useState } from 'react';
import axios from 'axios';
import {Breadcrumb, Table, Offcanvas, Button} from 'react-bootstrap';
import './DatasetExplorer.css';


interface IpniResult {
  Protocol: string;
  PeerID: string;
  Multiaddrs: string[];
}

interface Deal {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DealID: number | null;
  State: string;
  Client: string;
  ClientAddress: string;
  Provider: string;
  ProposalID: string;
  Label: string;
  PieceCID: string;
  PieceSize: number;
  Start: string;
  Duration: number;
  End: string;
  SectorStart: string;
  Price: number;
  Verified: boolean;
  ErrorMessage: string;
  ScheduleID: number | null;
}


interface Directory {
  ID: number;
  CID: string;
  Name: string;
  ParentID: number | null;
}

interface Item {
  ID: number;
  ScannedAt: string;
  ChunkID: number;
  Type: string;
  Path: string;
  Size: number;
  Offset: number;
  Length: number;
  LastModified: string | null;
  Version: number;
  CID: string;
  ErrorMessage: string;
  DirectoryID: number;
}

interface DatasetExplorerProps {
  rootDirectoryID: string;
}

const DatasetExplorer: React.FC<DatasetExplorerProps> = ({ rootDirectoryID }) => {
  const [directories, setDirectories] = useState<Directory[]>([]);
  const [items, setItems] = useState<Item[]>([]);
  const [breadcrumb, setBreadcrumb] = useState<Directory[]>([{ID: parseInt(rootDirectoryID), CID: "", Name: "", ParentID: null}]);
  const [selectedItem, setSelectedItem] = useState<Item | null>(null);
  const [showDeals, setShowDeals] = useState(false);

  useEffect(() => {
    fetchDirectoryEntries(breadcrumb);
  }, [breadcrumb]);

  const fetchDirectoryEntries = async (breadcrumb: Directory[]) => {
    try {
      const currentDirectory = breadcrumb[breadcrumb.length - 1];
      const response = await axios.get(`${process.env.REACT_APP_API_BASE_URL}/api/directory/${currentDirectory.ID}/entries`);
      setDirectories(response.data.Directories);
      setItems(response.data.Items);
    } catch (error) {
      console.error('Error fetching directory entries:', error);
    }
  };

  const handleDirectoryClick = (directoryID: string) => {
    const directory = directories.find((dir) => dir.ID.toString() === directoryID);
    setBreadcrumb([...breadcrumb, directory!]);
  };

  const handleBreadcrumbClick = (index: number) => {
    const newBreadcrumb = breadcrumb.slice(0, index + 1);
    setBreadcrumb(newBreadcrumb);
  };

  const handleItemClick = (itemID: string) => {
    const item = items.find((itm) => itm.ID.toString() === itemID);
    setSelectedItem(item!);
    setShowDeals(true);
  };

  return (
    <div className="dataset-explorer-container">
      <div className="dataset-explorer-main">
        <div>
          <Breadcrumb>
            {breadcrumb.map((dir, index) => (
              <Breadcrumb.Item key={dir.ID} onClick={() => handleBreadcrumbClick(index)}>
                {dir.Name}
              </Breadcrumb.Item>
            ))}
          </Breadcrumb>
          <Table striped bordered hover>
            <thead>
            <tr>
              <th>Name</th>
              <th>Size</th>
              <th>Offset</th>
              <th>Length</th>
              <th>CID</th>
            </tr>
            </thead>
            <tbody>
            {directories.map((dir) => (
              <tr key={dir.ID} onClick={() => handleDirectoryClick(dir.ID.toString())} style={{ cursor: 'pointer' }}>
                <td>{dir.Name}</td>
                <td></td>
                <td></td>
                <td></td>
                <td>{dir.CID}</td>
              </tr>
            ))}
            {items.map((item) => (
              <tr key={item.ID} onClick={() => handleItemClick(item.ID.toString())}>
                <td>{item.Path.split('/').pop()}</td>
                <td>{item.Size}</td>
                <td>{item.Offset}</td>
                <td>{item.Length}</td>
                <td>{item.CID}</td>
              </tr>
            ))}
            </tbody>
          </Table>
        </div>
      </div>
      {selectedItem && (
        <div className="deal-panel">
          <DealPanel item={selectedItem} show={showDeals} setShow={setShowDeals} />
        </div>
      )}
    </div>
  );
};

interface DealPanelProps {
  item: Item | null;
}

const DealPanel: React.FC<DealPanelProps & { show: boolean; setShow: (show: boolean) => void }> = ({
                                                                                                     item,
                                                                                                     show,
                                                                                                     setShow,
                                                                                                   }) => {
  const [deals, setDeals] = useState<Deal[]>([]);
  const [ipniResult, setIpniResult] = useState<IpniResult[]>([]);

  useEffect(() => {
    if (item) {
      fetchDeals(item.ID);
      axios.get(`https://cid.contact/cid/${item.CID}`).then((response) => {
        const newResult = [];
        for(const r1 of response.data.MultihashResults) {
          for (const r2 of r1.ProviderResults) {
            const peer = r2.Provider.ID;
            const protocol = "GraphSync";
            const multiaddrs = r2.Provider.Addrs;
            newResult.push( {Protocol: protocol, PeerID: peer, Multiaddrs: multiaddrs});
          }
        }
        setIpniResult(newResult.filter((item, index, self) => {
          return self.findIndex((obj) => JSON.stringify(obj) === JSON.stringify(item)) === index;
        }));
      });
    }
  }, [item]);

  const fetchDeals = async (itemID: number) => {
    try {
      const response = await axios.get(`${process.env.REACT_APP_API_BASE_URL}/api/item/${itemID}/deals`);
      const activeDeals = response.data.filter((deal: Deal) => deal.State === "active");
      setDeals(activeDeals);
    } catch (error) {
      console.error('Error fetching deals:', error);
    }
  };

  if (!item) {
    return null;
  }

  return (<Offcanvas
      show={show}
      onHide={() => setShow(false)}
      placement="end"
      className="deal-panel deal-panel-offcanvas"
    >
      <Offcanvas.Header closeButton>
        <Offcanvas.Title>Active Deals</Offcanvas.Title>
      </Offcanvas.Header>
      <Offcanvas.Body>
        <Table striped bordered hover size="sm">
          <thead>
          <tr>
            <th>Deal ID</th>
            <th>Provider</th>
            <th>Start</th>
            <th>End</th>
          </tr>
          </thead>
          <tbody>
          {deals.map((deal) => (
            <tr key={deal.ID}>
              <td>
                <a href={`https://filfox.info/en/deal/${deal.DealID}`} target="_blank" rel="noopener noreferrer">
                  {deal.DealID}
                </a>
              </td>
              <td>
                <a href={`https://filfox.info/en/address/${deal.Provider}`} target="_blank" rel="noopener noreferrer">
                  {deal.Provider}
                </a>
              </td>
              <td>{new Date(deal.Start).toLocaleDateString()}</td>
              <td>{new Date(deal.End).toLocaleDateString()}</td>
            </tr>
          ))}
          </tbody>
        </Table>
        <Table striped bordered hover size="sm">
          <thead>
          <tr>
            <th>Protocol</th>
            <th>Peer</th>
            <th>Addr</th>
          </tr>
          </thead>
          <tbody>
          {ipniResult.map((result) => (
            <tr key={result.PeerID}>
              <td>{result.Protocol}</td>
              <td>{result.PeerID}</td>
              <td>{result.Multiaddrs.join(", ")}</td>
            </tr>
          ))}
          </tbody>
        </Table>
        <Button variant="primary" >
          Download from Saturn
        </Button>
      </Offcanvas.Body>
    </Offcanvas>
  );
};


export default DatasetExplorer;
