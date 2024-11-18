import { useEffect, useState } from "react";
import {
  LineChart,
  Line,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import "./App.css";

function App() {
  const backendUrl = import.meta.env.BACKEND_ADDRESS ?? "http://pi.local:9999";
  const [energyData, setEnergyData] = useState([]);
  const [date, setDate] = useState(new Date());

  useEffect(() => {
    const currentDate = new Date().toISOString().split("T")[0];
    const fetchData = async () => {
      const response = await fetch(`${backendUrl}/energy/${currentDate}`);
      const json = await response.json();
      setEnergyData(json);
    };
    fetchData();
  }, []);

  console.log(energyData);

  return (
    <main>
      <h1>Prices for the day: {date.toLocaleDateString("de-DE")}</h1>
      <ResponsiveContainer width="100%" height={400}>
        <LineChart data={energyData}>
          <Line type="monotone" dataKey="value" stroke="#8884d8" />
          <CartesianGrid stroke="#ccc" />
          <XAxis
            dataKey="date"
            tickFormatter={(value) => {
              const date = new Date(value);
              return date.toLocaleTimeString("de-DE", {
                hour: "2-digit",
                minute: "2-digit",
              });
            }}
          />
          <YAxis />
          <Tooltip
            formatter={(value, name, props) => {
              return [`${value} ct/kWh`, "Price"];
            }}
            labelFormatter={(label) => {
              const date = new Date(label);
              return date.toLocaleTimeString("de-DE", {
                hour: "2-digit",
                minute: "2-digit",
              });
            }}
          />
        </LineChart>
      </ResponsiveContainer>
    </main>
  );
}

export default App;
