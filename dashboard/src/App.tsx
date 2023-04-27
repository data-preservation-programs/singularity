import React, { useState } from 'react';
import { Container, Navbar, Row, Col } from 'react-bootstrap';
import DatasetList from './components/DatasetList';
import OverallDealStats from './components/OverallDealStats';
import DatasetDealStats from './components/DatasetDealStats';

const App: React.FC = () => {
  const [selectedDatasetId, setSelectedDatasetId] = useState<string | null>(null);

  const handleBrandClick = () => {
    window.location.reload();
  };

  const handleDatasetClick = (id: string) => {
    setSelectedDatasetId(id);
  };

  return (
    <div>
      <Navbar bg="dark" variant="dark">
        <Navbar.Brand href="#" onClick={handleBrandClick}>
          Singularity
        </Navbar.Brand>
      </Navbar>
      <Container fluid>
        <Row>
          <Col md={4} lg={3}>
            <DatasetList onDatasetClick={handleDatasetClick} />
          </Col>
          <Col md={8} lg={9}>
            {selectedDatasetId ? (
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
