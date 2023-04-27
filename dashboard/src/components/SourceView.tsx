import React from 'react';
import { Tab, Tabs } from 'react-bootstrap';

interface SourceViewProps {
  sourceID: string;
}

const SourceView: React.FC<SourceViewProps> = ({ sourceID }) => {
  return (
    <Tabs defaultActiveKey="datasetExplorer" id="source-view-tabs">
      <Tab eventKey="datasetExplorer" title="Dataset Explorer">
        {/* Add Dataset Explorer content here */}
        <p>Dataset Explorer content for source {sourceID} goes here.</p>
      </Tab>
      <Tab eventKey="carExplorer" title="Car Explorer">
        {/* Add Car Explorer content here */}
        <p>Car Explorer content for source {sourceID} goes here.</p>
      </Tab>
    </Tabs>
  );
};

export default SourceView;
