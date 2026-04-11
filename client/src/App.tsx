import { useEffect, useState } from "react"

import "./App.css"

function App() {
    const [message, setMessage] = useState("")
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        // Fetch data from the backend API

        const apiUrl =
            process.env.NODE_ENV === "production"
                ? "/api/status"
                : "http://localhost:3002/api/status"

        console.log("Fetching from:", apiUrl)

        fetch(apiUrl)
            .then((response) => {
                console.log("Response status:", response.status)
                return response.json()
            })
            .then((data) => {
                console.log("Data received:", data)
                setMessage(data.status)
                setLoading(false)
            })
            .catch((error) => {
                console.error("Error fetching data:", error)
                setMessage("Error connecting to server")
                setLoading(false)
            })
    }, [])

    return (
        <div className="App">
            <header className="App-header">
                <h1>React + Go with Gin + amir + abc</h1>
                <p>{loading ? "Loading..." : `API Response: ${message}`}</p>
                <p className="environment">
                    Environment: {process.env.NODE_ENV}
                </p>
            </header>
        </div>
    )
}

export default App
