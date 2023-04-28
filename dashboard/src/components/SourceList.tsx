import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Button } from 'react-bootstrap';

interface Source {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DatasetID: number;
  Type: string;
  Path: string;
  ScanInterval: number;
  ScanningState: string;
  ScanningWorkerID: string | null;
  LastScanned: string;
  MaxWait: number;
  ErrorMessage: string;
  RootDirectoryID: number;
}

interface SourceListProps {
  datasetID: number;
  onSourceClick: (sourceID: string, rootDirectoryID: string) => void;
}
const SourceList: React.FC<SourceListProps> = ({ datasetID, onSourceClick }) => {
  const [sources, setSources] = useState<Source[]>([]);
  const [fetched, setFetched] = useState(false);

  useEffect(() => {
    if (!fetched) {
      fetchSources();
    }
  }, [datasetID, fetched]);

  const fetchSources = async () => {
    try {
      const response = await axios.get<Source[]>(`${process.env.REACT_APP_API_BASE_URL}/api/dataset/${datasetID}/sources`);
      setSources(response.data);
      setFetched(true);
    } catch (error) {
      console.error('Error fetching sources:', error);
    }
  };

  return (
    <div className="list-group">
      {sources.map((source) => (
        <Button
          key={source.ID}
          variant="outline-secondary"
          className="list-group-item list-group-item-action"
          onClick={() => onSourceClick(source.ID.toString(), source.RootDirectoryID.toString())}
        >
          {source.Path}
        </Button>
      ))}
    </div>
  );
};

export default SourceList;
