import React from 'react';
import {client} from "./client";

const App: React.FC = () => {
    const [content, setContent] = React.useState<[string, string][]>([])
    React.useEffect(() => {
        const fetchData = async () => {
            let content: [string, string][] = []

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
                        <pre><code>{data}</code></pre>
                    </div>
                )
            )}
        </div>
    );
};

export default App;
