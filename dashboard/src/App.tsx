import React from 'react';
import {client} from "./client";
import {CodeBlock, dracula} from "react-code-blocks";

const App: React.FC = () => {
    const [content, setContent] = React.useState<[string, string][]>([])
    React.useEffect(() => {
        const fetchData = async() => {
            let content:[string, string][] = []

            const preparations = await client.preparation.listPreparations()
            content.push(['preparations', JSON.stringify(preparations.data, null, 2)])

            setContent(content)
        }
        fetchData().catch(console.error)
    }, [])
    return (
        <div>
            {content.map(([title, data]) => (
                <div>
                    <h2>{title}</h2>
                    <CodeBlock
                      text={data}
                      language={'json'}
                      showLineNumbers={true}
                      theme={dracula}
                    />
                </div>
            )
            )}
        </div>
    );
};

export default App;
