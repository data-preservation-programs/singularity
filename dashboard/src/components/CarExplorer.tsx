import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Table, Offcanvas, Pagination } from 'react-bootstrap';
import './DatasetExplorer.css';

interface Car {
  ID: number;
  CreatedAt: string;
  PieceCID: string;
  PieceSize: number;
  RootCID: string;
  FileSize: number;
  FilePath: string;
}

interface CarExplorerProps {
  sourceID: string;
}

const CarExplorer: React.FC<CarExplorerProps> = ({ sourceID }) => {
  const [cars, setCars] = useState<Car[]>([]);
  const [selectedCar, setSelectedCar] = useState<Car | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 25;
  const [showOffcanvas, setShowOffcanvas] = useState(false);


  useEffect(() => {
    fetchCars(sourceID);
  }, [sourceID]);

  const handlePageChange = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  const renderPagination = () => {
    const totalPages = Math.ceil(cars.length / itemsPerPage);

    let items = [];
    for (let number = 1; number <= totalPages; number++) {
      items.push(
        <Pagination.Item
          key={number}
          active={number === currentPage}
          onClick={() => handlePageChange(number)}
        >
          {number}
        </Pagination.Item>
      );
    }

    return <Pagination>{items}</Pagination>;
  };


  const fetchCars = async (sourceID: string) => {
    try {
      const response = await axios.get(`${process.env.REACT_APP_API_BASE_URL}/api/source/${sourceID}/cars`);
      setCars(response.data);
    } catch (error) {
      console.error('Error fetching cars:', error);
    }
  };
  const handleCarClick = (carID: string) => {
    const car = cars.find((c) => c.ID.toString() === carID);
    setSelectedCar(car!);
    setShowOffcanvas(true); // Add this line
  };

  return (
    <div>
      {renderPagination()}
      <Table striped bordered hover>
        <thead>
        <tr>
          <th>Piece CID</th>
          <th>Piece Size</th>
          <th>Root CID</th>
          <th>File Size</th>
        </tr>
        </thead>
        <tbody>
        {cars
          .slice((currentPage - 1) * itemsPerPage, currentPage * itemsPerPage)
          .map((car) => (
            <tr key={car.ID} onClick={() => handleCarClick(car.ID.toString())} style={{ cursor: 'pointer' }}>
              <td>{car.PieceCID}</td>
              <td>{car.PieceSize}</td>
              <td>{car.RootCID}</td>
              <td>{car.FileSize}</td>
            </tr>
          ))}
        </tbody>
      </Table>
      {selectedCar &&
          <div className="deal-panel">
              <CarOffCanvas
                  car={selectedCar}
                  show={showOffcanvas}
                  onHide={() => setShowOffcanvas(false)}
              />
          </div>}
    </div>
  );
};

// Add your deal, item and CarOffCanvasProps interfaces here

// Add your Deal, Item, and CarOffCanvasProps interfaces here


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

interface CarOffCanvasProps {
  car: Car | null;
  show: boolean;
  onHide: () => void;
}

const CarOffCanvas: React.FC<CarOffCanvasProps> = ({ car, show, onHide }) => {
  const [deals, setDeals] = useState<Deal[]>([]);
  const [items, setItems] = useState<Item[]>([]);

  useEffect(() => {
    if (car) {
      fetchDeals(car.ID);
      fetchItems(car.ID);
    }
  }, [car]);

  const fetchDeals = async (carID: number) => {
    try {
      const response = await axios.get(`${process.env.REACT_APP_API_BASE_URL}/api/car/${carID}/deals`);
      setDeals(response.data.filter((d: Deal) => d.State == 'active'));
    } catch (error) {
      console.error('Error fetching deals:', error);
    }
  };

  const fetchItems = async (carID: number) => {
    try {
      const response = await axios.get(`${process.env.REACT_APP_API_BASE_URL}/api/car/${carID}/items`);
      setItems(response.data);
    } catch (error) {
      console.error('Error fetching items:', error);
    }
  };

  return (
    <Offcanvas show={show} onHide={onHide} placement="end" className="car-offcanvas deal-panel deal-panel-offcanvas">
      <Offcanvas.Header closeButton>
        <Offcanvas.Title>Car Details</Offcanvas.Title>
      </Offcanvas.Header>
      <Offcanvas.Body>
        <h5>Deals</h5>
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
              <a href={`https://filfox.info/en/deal/${deal.DealID}`} target="_blank" rel="noreferrer">
                {deal.DealID}
              </a>
            </td>
            <td>
              <a href={`https://filfox.info/en/address/${deal.Provider}`} target="_blank" rel="noreferrer">
                {deal.Provider}
              </a>
            </td>
            <td>{new Date(deal.Start).toLocaleDateString()}</td>
            <td>{new Date(deal.End).toLocaleDateString()}</td>
          </tr>
        ))}
        </tbody>
      </Table>
        <h5>Items</h5>
        <Table striped bordered hover size="sm">
        <thead>
        <tr>
          <th>Path</th>
          <th>CID</th>
          <th>Length</th>
        </tr>
        </thead>
        <tbody>
        {items.map((item) => (
          <tr key={item.ID}>
            <td>{item.Path}</td>
            <td>{item.CID}</td>
            <td>{item.Length}</td>
          </tr>
        ))}
        </tbody>
      </Table>
      </Offcanvas.Body>
    </Offcanvas>
  );
};


export default CarExplorer;
