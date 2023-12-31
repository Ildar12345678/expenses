import React, {useContext, useEffect, useState} from "react";
import StatisticChart from "./StatisticChart";
import ExpensesContext from "./Context";

function GetStatistic() {
    // const { stats } = useContext(ExpensesContext);
    const { random } = useContext(ExpensesContext);

    const [stats, setStats] = useState([]);

    useEffect(() => {
        fetch('http://localhost:4000/stats')
            .then(response => response.json())
            .then(data => {
                // Sort the data based on a specific property
                // const sortedData = data.sort((a, b) => b.SumSubcat - a.SumSubcat);

                // Update the state with the sorted data
                setStats(data);
            })
            .catch(console.log);

    }, [random]);

    if (stats) {
        let sum = 0
        for (let el of stats) {
            sum += el.SumSubcat
        }
        return (
            <>
                <div style={{ display: 'flex', alignItems: 'center' }}>
                    <div className="main-container">
                        <ul>
                            <div className="box">
                                Sum = {sum}
                            </div>
                            {stats.map((stat, i) =>
                                (
                                    <div key={i} className="container">
                                        <div className="box">{stat.Cat}</div>
                                        <div className="box">{stat.Subcat}</div>
                                        <div className="box">{stat.SumSubcat}</div>
                                    </div>
                                ))}
                        </ul>
                    </div>
                    <StatisticChart data={stats} />
                </div>
            </>
        );
    }

    return null;
}

export default function StatisticContainer() {
    return <GetStatistic />;
}