import React, { useState } from "react";
import Sidenav from "../../Layouts/Sidenav";
import { useLocation, useNavigate } from "react-router-dom";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import AddIcon from "@mui/icons-material/Add";
import DeleteIcon from "@mui/icons-material/Delete";
import SaveIcon from "@mui/icons-material/Save";
import { Button } from "@mui/material";
import { useFormik } from "formik";

export default function BarangAdd() {
  const navigate = useNavigate();
  const { state } = useLocation();
  var modelData = {
    Id: 0,
    Code: "",
    Name: "",
    Total: 0,
    Description: "",
    IsActive: 1,
  };

  const [arrayData, setArrayData] = useState([]);

  const formik = useFormik({
    initialValues: {
      modelData,
    },
  });

  if (state === null) {
    // setArrayData([...arrayData, "modelData"]);
  }

  const handleInputChange = (e, index) => {
    formik.setFieldValue(e.target.name, e.target.value);
  };

  const addData = () => {
    setArrayData([...arrayData, formik]);
  };

  const deleteData = (key) => {
    const filtered = arrayData.filter((item, index) => index !== key);
    setArrayData(filtered);
  };

  return (
    <Sidenav
      children={
        <div>
          <h1>Barang</h1>
          <div className="my-3">
            <Button
              variant="contained"
              color="secondary"
              onClick={() => {
                navigate("/barang");
              }}
            >
              <ArrowBackIcon /> Back
            </Button>
            <Button
              className="ms-3"
              variant="contained"
              color="primary"
              onClick={addData}
            >
              <AddIcon /> Add
            </Button>
            <Button variant="contained" color="success" className="ms-3">
              <SaveIcon /> Save
            </Button>
          </div>
          {arrayData.map((row, key) => (
            <div key={key} className="card p-3">
              <Button
                className="ms-3"
                variant="contained"
                color="primary"
                onClick={() => deleteData(key)}
              >
                <DeleteIcon /> Delete
              </Button>
              <div className="mb-3">
                <label className="form-label">Code</label>
                <input
                  onChange={(e) => handleInputChange(e, key)}
                  type="text"
                  className="form-control"
                  placeholder="Code"
                  name="Code"
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Name</label>
                <input
                  onChange={(e) => handleInputChange(e, key)}
                  type="text"
                  className="form-control"
                  placeholder="Name"
                  name="Name"
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Total</label>
                <input
                  onChange={(e) => handleInputChange(e, key)}
                  type="number"
                  className="form-control"
                  placeholder="Total"
                  name="Total"
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Description</label>
                <input
                  onChange={(e) => handleInputChange(e, key)}
                  type="text"
                  className="form-control"
                  placeholder="Description"
                  name="Description"
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Status</label>
                <select
                  onChange={(e) => handleInputChange(e, key)}
                  name="IsActive"
                  className="form-select"
                >
                  <option value="1">Active</option>
                  <option value="0">Inactive</option>
                </select>
              </div>
            </div>
            // <div key={index} className="card p-3"></div>
          ))}
        </div>
      }
    />
  );
}
