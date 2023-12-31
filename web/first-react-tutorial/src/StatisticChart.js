import React, {useEffect, useState, useRef} from "react"
import * as d3 from 'd3'
import {schemeSet1, schemeSet2, schemeSet3} from 'd3-scale-chromatic'

const PieChart = ({ data }) => {
    const chartRef = useRef();
    const chartRef2 = useRef();
    const [tooltip, setTooltip] = useState(null);
    useEffect(() => {
        let dataInsideSubcats = [
            {Cat:'здоровье',SumSubcat:0},
            {Cat:'непродукты',SumSubcat:0},
            {Cat:'продукты',SumSubcat:0},
            {Cat:'проезд',SumSubcat:0},
            {Cat:'развлечения',SumSubcat:0},
            {Cat:'социальные',SumSubcat:0},
        ]
        if (data) {
            let toAddHealth = 0
            let toAddProd = 0
            let toAddNoProd = 0
            let toAddProezd = 0
            let toAddFunny = 0
            let toAddOther = 0

             let func = (data) => {
                for (const el of data) {
                    switch (el.Cat) {
                        case 'здоровье':
                            toAddHealth += el.SumSubcat
                            break
                        case 'непродукты':
                            toAddNoProd += el.SumSubcat
                            break
                        case 'продукты':
                            toAddProd += el.SumSubcat
                            break
                        case 'проезд':
                            toAddProezd += el.SumSubcat
                            break
                        case 'развлечения':
                            toAddFunny += el.SumSubcat
                            break
                        case 'социальные':
                            toAddOther += el.SumSubcat
                            break
                    }
                }
            }
            func(data)
            for (const el of dataInsideSubcats) {
                switch (el.Cat) {
                    case 'здоровье':
                        el.SumSubcat = toAddHealth
                        break
                    case 'непродукты':
                        el.SumSubcat = toAddNoProd
                        break
                    case 'продукты':
                        el.SumSubcat = toAddProd
                        break
                    case 'проезд':
                        el.SumSubcat = toAddProezd
                        break
                    case 'развлечения':
                        el.SumSubcat = toAddFunny
                        break
                    case 'социальные':
                        el.SumSubcat = toAddOther
                        break
                }
            }
            drawChart(chartRef, data);
            drawChart(chartRef2, dataInsideSubcats);
        }
    }, [data]);

    const drawChart = (ref, dataToShow) => {
        const svg = d3.select(ref.current);
        const width = +svg.attr('width');
        const height = +svg.attr('height');
        const radius = Math.min(width, height) / 2;

        const pie = d3.pie().value(d => d.SumSubcat);
        const arc = d3.arc().innerRadius(0).outerRadius(radius);

        const color = d3.scaleOrdinal(schemeSet3);

        const arcs = pie(dataToShow);

        const arcWithTooltip = d3
            .arc()
            .innerRadius(0)
            .outerRadius(radius)
            .cornerRadius(10)
            .padAngle(0.01);

        svg
            .selectAll('path')
            .data(arcs)
            .join('path')
            .attr('d', arcWithTooltip)
            .attr('fill', (d, i) => color(i))
            .attr('transform', `translate(${width / 2}, ${height / 2})`)
            .on('mouseover', (event, d) => {
                setTooltip({
                    data: d,
                    x: 100,
                    y: event.pageY,
                    style: 'bisque'
                });
            })
            .on('mouseout', () => {
                setTooltip(null);
            });

        svg
            .selectAll('text')
            .data(arcs)
            .join('text')
            .attr('transform', d => `translate(${arc.centroid(d)})`)
            .attr('dy', '0.35em')
            .text(d => d.data.label)
            .style('text-anchor', 'middle');
    };

    return (

        <div style={{ }}>

            <div style={{display: 'flex', width: '100%', height: '100%' }}>
                <svg ref={chartRef}
                     viewBox={`0 0 500 500`} // Adjust the viewBox dimensions accordingly
                     width={500}
                     height={500}
                     style={{ marginLeft: 'auto', marginRight: 'auto', display: 'block' }}
                >
                </svg>
                <svg ref={chartRef2}
                     viewBox={`0 0 600 600`} // Adjust the viewBox dimensions accordingly
                     width={500}
                     height={500}
                     style={{ marginLeft: 'auto', marginRight: 'auto', display: 'block' }}
                >
                </svg>
                <div className="label" >
                    {tooltip && (
                        <div style={{backgroundColor: tooltip.style}}>
                            {/* Render the tooltip or information about the hovered sector */}
                            <p>{tooltip.data.data.Cat}</p>
                            <p>{tooltip.data.data.Subcat}</p>
                            <p>{tooltip.data.data.SumSubcat}</p>
                        </div>
                    )}
                </div>

            </div>

        </div>
    );
};

export default PieChart;