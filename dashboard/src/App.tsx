import React, {useState} from 'react';
import {Col, Container, Navbar, Row} from 'react-bootstrap';
import DatasetList from './components/DatasetList';
import OverallDealStats from './components/OverallDealStats';
import DatasetDealStats from './components/DatasetDealStats';
import SourceView from "./components/SourceView";
import './App.css';

const App: React.FC = () => {
  const [selectedDatasetId, setSelectedDatasetId] = useState<string | null>(null);
  const [selectedSourceID, setSelectedSourceID] = useState<string | null>(null);
  const [selectedRootDirectoryId, setSelectedRootDirectoryId] = useState<string | null>(null);
  const handleBrandClick = () => {
    window.location.reload();
  };

  const handleDatasetClick = (id: string) => {
    setSelectedDatasetId(id);
    setSelectedSourceID(null);
  };
  const handleSourceClick = (sourceID: string, rootDirectoryID: string) => {
    setSelectedDatasetId(null);
    setSelectedSourceID(sourceID);
    setSelectedRootDirectoryId(rootDirectoryID);
  };

  return (
    <div>
      <Navbar bg="dark" variant="dark">
        <Navbar.Brand href="#" onClick={handleBrandClick}>
          Singularity Dashboard
        </Navbar.Brand>
      </Navbar>
      <Container fluid>
        <Row>
          <Col md={4} lg={3}>
            <DatasetList onDatasetClick={handleDatasetClick} onSourceClick={handleSourceClick} />
          </Col>
          <Col md={8} lg={9}>
            {selectedSourceID ? (
              <SourceView sourceID={selectedSourceID.toString()} rootDirectoryID={selectedRootDirectoryId!} />
            ) : selectedDatasetId ? (
              <DatasetDealStats datasetId={selectedDatasetId} />
            ) : (
              <OverallDealStats />
            )}
          </Col>
        </Row>
      </Container>
    </div>
  );
};

export default App;
