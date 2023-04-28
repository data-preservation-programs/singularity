import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Accordion, Card, Button } from 'react-bootstrap';
import SourceList from './SourceList';

interface Dataset {
  ID: number;
  Name: string;
  CreatedAt: string;
  UpdatedAt: string;
  MinSize: number;
  MaxSize: number;
  PieceSize: number;
  OutputDirs: string[];
  EncryptionRecipients: string[];
  EncryptionScript: string;
  Wallets: Wallet[];
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
    <Accordion defaultActiveKey="0">
      {datasets.map((dataset, index) => (
        <Card key={dataset.ID}>
          <Card.Header>
            <Button
              variant="link"
              onClick={() => {
                onDatasetClick(dataset.ID.toString());
                document.getElementById(`dataset-${dataset.ID}`)?.classList.toggle('show');
              }}
              aria-controls={`dataset-${dataset.ID}`}
              aria-expanded="false"
            >
              {dataset.Name}
            </Button>
          </Card.Header>
          <div id={`dataset-${dataset.ID}`} className="collapse" data-parent="#accordion">
            <Card.Body>
              <SourceList datasetID={dataset.ID} onSourceClick={onSourceClick} />
            </Card.Body>
          </div>
        </Card>
      ))}
    </Accordion>
  );
};

export default DatasetList;
