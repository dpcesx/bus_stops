import React from "react"
import {createRoot} from "react-dom/client"
import { DateTime } from "luxon"

const getArrivalEstimate = async (stopNumber) => {
    const response = await fetch(`http://localhost:8080/stop/${stopNumber}`)
    return await response.json()    
}
const App = () => {
    const [now, setNow] = React.useState(DateTime.local())
    const [estimates, setEstimates] = React.useState([])
    const effect = () => {
        (async () => {
            const estimates = [await getArrivalEstimate(1), await getArrivalEstimate(2)]
            setEstimates(estimates)
            console.log(estimates)
        })()

        const intervalHandler = () => {
            const newNow = DateTime.local()
            setNow(newNow)
            
        }
        const intervalId = setInterval(intervalHandler, 1000)

        return () => clearInterval(intervalId)
    }
    const deps = [setInterval, clearInterval, setNow, DateTime.local]

    React.useEffect(effect, deps)

    return (
        <>
            <h1>{now.toLocaleString(DateTime.TIME_SIMPLE)}</h1>
            {estimates.map((stops, idx) => (
                    <>
                        <h2>Stop {idx + 1}</h2>
                        {stops.map((routes, jdx) => (
                                <>
                                    <h3>Route {jdx + 1}</h3>
                                    <p key={idx}>{routes.join(" mins ")} mins</p>
                                </>
                            )
                        )}
                    </>
                )
            )}
            {/* <pre>{JSON.stringify(estimates, null, 2)}</pre> */}
        </>
    )
}

const container = document.getElementById('app')
const root = createRoot(container)
root.render(<App />)