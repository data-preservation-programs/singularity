import React from 'react';
import { Tab, Tabs } from 'react-bootstrap';
import DatasetExplorer from "./DatasetExplorer";
import CarExplorer from "./CarExplorer";

interface SourceViewProps {
  sourceID: string;
  rootDirectoryID: string;
}

const SourceView: React.FC<SourceViewProps> = ({ sourceID, rootDirectoryID  }) => {
  return (
    <Tabs defaultActiveKey="datasetExplorer" id="source-view-tabs">
      <Tab eventKey="datasetExplorer" title="Dataset Explorer">
        <DatasetExplorer rootDirectoryID={rootDirectoryID} />
      </Tab>
      <Tab eventKey="carExplorer" title="Car Explorer">
        <CarExplorer sourceID={sourceID} />
      </Tab>
    </Tabs>
  );
};

export default SourceView;
