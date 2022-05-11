import "./App.css";
import { Grid, TextField, Button } from "@mui/material";
import { useState } from "react";

function App() {
  let [title, setTitle] = useState("");
  let [content, setContent] = useState("");

  function submit() {
    console.log(title + "  " + content);
  }

  return (
    <div className="App">
      <Grid container spacing={2} style={{ height:"auto" }}>
        <Grid item style={{ textAlign: "center", width: "100%" }}>
          <TextField onChange={(e) => { setTitle(e.target.value) }} value={title} label="Title" style={{ width:"500px" }} />
        </Grid>
        <Grid item style={{ textAlign: "center", width: "100%", height:"200px" }}>
          <TextField onChange={(e) => { setContent(e.target.value) }} value={content} id="outlined-multiline-static" label="Content" multiline rows={6} style={{ width:"500px" }} />
        </Grid>
        <Grid item style={{ textAlign: "center", width: "100%" }}>
          <Button onClick={submit} variant="contained">Commit</Button>
        </Grid>
      </Grid>
    </div>
  );
}

export default App;
