import { useEffect, useState } from "react"
import FileInput from "./components/FileInput.jsx"
import DropdownList from "./components/DropdownList.jsx"
import Papa from "papaparse"
import Input from "./components/Input.jsx"
import Button from "./components/Button.jsx"

const initialInputs = {
  host: "jqii6ig8f6.asia-southeast1.gcp.clickhouse.cloud",
  port: "9440",
  database: "default",
  username: "default",
  password: "8ee_Rt6D~CWK7",
  tableName: ""
}

function App() {
  const [columns, setColumns] = useState([])
  const [file, setFile] = useState(null)
  const [selectedColumns, setSelectedColumns] = useState({})
  const [inputs, setInputs] = useState(initialInputs)

  useEffect(() => {
    getColumns(file, initializeColumns, setColumns, setSelectedColumns)
  }, [file])

  useEffect(() => {
    console.log("Selected columns")
    console.log(selectedColumns)
  }, [selectedColumns])


  const handleFileChange = (event) => {
    const file = event.target.files[0]
    setFile(file)
  }

  const handleColumnSelection = (event) => {
    const id = event.target.id
    const checked = event.target.checked

    console.log("Marking: " + id + " as " + checked)

    setSelectedColumns(selectedColumns => {
      return { ...selectedColumns, [`${id}`]: checked }
    })
  }

  const handleInputChange = (event) => {
    const id = event.target.id
    const value = event.target.value

    setInputs(inputs => {
      return { ...inputs, [id]: value }
    })
  }

  return (
    <div>
      <FileInput
        file={file}
        onChange={handleFileChange}
      />

      <DropdownList
        items={columns}
        onChange={handleColumnSelection}
      />

      <Input
        placeholder={"Hostname"}
        value={inputs.host}
        id={"host"}
        onChange={handleInputChange}
      >
      </Input>

      <Input
        placeholder={"Port"}
        value={inputs.port}
        id={"port"}
        onChange={handleInputChange}
      ></Input>

      <Input
        placeholder={"Database"}
        value={inputs.database}
        id={"database"}
        onChange={handleInputChange}
      ></Input>

      <Input
        placeholder={"Table Name"}
        value={inputs.tableName}
        id={"tableName"}
        onChange={handleInputChange}
      ></Input>

      <Input
        placeholder={"Username"}
        value={inputs.username}
        id={"username"}
        onChange={handleInputChange}
      ></Input>

      <Input
        placeholder={"Password"}
        value={inputs.password}
        id={"password"}
        onChange={handleInputChange}
      ></Input>

      <Button
        text={"Connect"}
        onClick={() => connect(inputs)}
      ></Button>

      <Button
        text={"Upload"}
        onClick={() => upload(columns, selectedColumns, inputs, file)}
      ></Button>
    </div>
  )
}

function getColumns(file, onComplete, setColumns, setSelectedColumns) {
  if (!file) {
    return
  }

  console.log("Parsing")

  Papa.parse(file,
    {
      preview: 1,
      header: true,
      complete: (results) => {
        console.log(results.meta.fields)
        onComplete(results.meta.fields, setColumns, setSelectedColumns)
      }
    }
  )
}

function initializeColumns(columns, setColumns, setSelectedColumns) {
  setColumns(columns)

  let tempColumnObjects = {}

  for (const column of columns) {
    tempColumnObjects[column] = true
  }

  setSelectedColumns(tempColumnObjects)
}

async function connect(inputs) {
  console.log(inputs)

  const data = {
    Host: inputs.host,
    Port: parseInt(inputs.port),
    Database: inputs.database,
    Username: inputs.username,
    Password: inputs.password
  }

  console.log(data)

  const res = await fetch('http://localhost:8080/connect', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json'
    }
  })

  const resJson = await res.json()

  console.log(resJson)
}


async function upload(columns, selectedColumnsObj, inputs, file) {
  const queue = []
  let processing = false

  const selectedColumns = Object.entries(selectedColumnsObj)
    .filter(([_, checked]) => checked)
    .map(([key]) => key)

  console.log("Selectec cols: ")
  console.log(selectedColumns)

  const uploadChunk = async function () {
    console.log("Trying")
    if (queue.length == 0) {
      processing = false
      return
    }

    processing = true
    const chunkInfo = queue.shift()

    const chunkNumber = chunkInfo.chunkNumber
    const chunk = chunkInfo.chunk

    console.log(chunkInfo)
    console.log("Processing chunk number: " + chunkNumber)
    const data = {
      ConnectionInfo: {
        Host: inputs.host,
        Port: parseInt(inputs.port),
        Database: inputs.database,
        Username: inputs.username,
        Password: inputs.password
      },
      TableName: inputs.tableName,
      ColumnNames: selectedColumns,
      Rows: chunk
    }

    const res = await fetch('http://localhost:8080/post', {
      method: 'POST',
      body: JSON.stringify(data),
      headers: {
        'Content-Type': 'application/json'
      }
    })

    console.log("Got chunk number: " + chunkNumber)

    const resJson = await res.json()

    console.log(resJson)
    uploadChunk()
  }

  let chunkNumber = 0

  Papa.parse(file, {
    header: false,
    dynamicTyping: true,
    skipEmptyLines: true,
    worker: true,
    chunkSize: 5000000,
    chunk: async (results, parser) => {
      let data

      if (chunkNumber == 0) {
        data = results.data.slice(1)
      } else {
        data = results.data
      }
      chunkNumber++
      console.log("Chunk:")
      console.log(data)

      const selectedIndices = selectedColumns.map(col => {
        const index = columns.indexOf(col)
        if (index !== -1) {
          return index
        }
      })

      const filteredData = data.map(row => {
        return selectedIndices.map(idx => row[idx])
      })

      console.log(filteredData)

      queue.push({ chunkNumber: chunkNumber, chunk: filteredData })

      if (processing == false) {
        uploadChunk()
      }
    },
    complete: () => {
      if (processing == false) {
        uploadChunk()
      }
      console.log("CSV upload complete.")
    }
  })
}

export default App
