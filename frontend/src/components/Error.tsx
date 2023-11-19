import { useRouteError } from "react-router-dom"

export default function Error(){
    const error = useRouteError()
    console.log(error)
    return <h2>Error while fetching data</h2>
}