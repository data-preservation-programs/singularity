import React from 'react';
import {client} from "./client";

const App: React.FC = () => {
    const [content, setContent] = React.useState<[string, string][]>([])
    React.useEffect(() => {
        const fetchData = async () => {
            let content: [string, string][] = []

            const preparations = await client.preparation.listPreparations()
            content.push(['preparations', JSON.stringify(preparations.data, null, 2)])

            const storages = await client.storage.listStorages()
            content.push(['storages', JSON.stringify(storages.data, null, 2)])

            const wallets = await client.wallet.listWallets()
            content.push(['wallets', JSON.stringify(wallets.data, null, 2)])

            const schedules = await client.schedule.listSchedules()
            content.push(['schedules', JSON.stringify(schedules.data, null, 2)])

            const pieces = await client.preparation.listPieces("1")
            content.push(['pieces', JSON.stringify(pieces.data, null, 2)])

            const rootEntries = await client.preparation.explorePreparation("1", "1", "/")
            content.push(['root directory', JSON.stringify(rootEntries.data, null, 2)])

            const subFolderEntries = await client.preparation.explorePreparation("1", "1", "/large_files")
            content.push(['sub directory', JSON.stringify(subFolderEntries.data, null, 2)])

            const fileDetail = await client.file.getFile(334)
            content.push(['file detail', JSON.stringify(fileDetail.data, null, 2)])

            const fileDeals = await client.file.getFileDeals(334)
            content.push(['file deals', JSON.stringify(fileDeals.data, null, 2)])

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
