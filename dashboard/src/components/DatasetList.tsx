import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Accordion, Card, Button } from 'react-bootstrap';
import SourceList from './SourceList';
import './DatasetList.css';

interface Dataset {
  id: number;
  name: string;
}

interface Wallet {
  ID: number;
  Name: string;
  CreatedAt: string;
  UpdatedAt: string;
  Address: string;
  TokenType: string;
  Default: boolean;
  PublicKey: string;
  PrivateKey: string;
}

interface DatasetListProps {
  onDatasetClick: (id: string) => void;
  onSourceClick: (sourceID: string, rootDirectoryID: string) => void;
}

const DatasetList: React.FC<DatasetListProps> = ({ onDatasetClick, onSourceClick }) => {
  const [datasets, setDatasets] = useState<Dataset[]>([]);

  useEffect(() => {
    fetchDatasets();
  }, []);

  const fetchDatasets = async () => {
    try {
      const response = await axios.get<Dataset[]>(`${process.env.REACT_APP_API_BASE_URL}/api/datasets`);
      setDatasets(response.data);
    } catch (error) {
      console.error('Error fetching datasets:', error);
    }
  };


  return (
    <Accordion defaultActiveKey="0" className="dataset-list">
      {datasets.map((dataset, index) => (
        <Card key={dataset.id}>
          <Card.Header>
            <Button className="dataset-button"
              variant="link"
              onClick={() => {
                onDatasetClick(dataset.id.toString());
                document.getElementById(`dataset-${dataset.id}`)?.classList.toggle('show');
              }}
              aria-controls={`dataset-${dataset.name}`}
              aria-expanded="false"
            >
              {dataset.name}
            </Button>
          </Card.Header>
          <div id={`dataset-${dataset.id}`} className="collapse" data-parent="#accordion">
            <Card.Body>
              <SourceList datasetID={dataset.id} onSourceClick={onSourceClick} />
            </Card.Body>
          </div>
        </Card>
      ))}
    </Accordion>
  );
};

export default DatasetList;
