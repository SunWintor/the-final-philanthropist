
lineConfig = {
  type: 'line',
  data: {
    labels: chartsLabels,
    datasets: [
      {
        label: 'Organic',
        /**
         * These colors come from Tailwind CSS palette
         * https://tailwindcss.com/docs/customizing-colors/#default-color-palette
         */
        backgroundColor: '#fef9c3',
        borderColor: '#fef9c3',
        data: [43, 48, 40, 108, 0, 0, 70],
        fill: false,
      },
      {
        label: 'Paid',
        fill: false,
        backgroundColor: '#7e3af2',
        borderColor: '#7e3af2',
        data: [24, 50, 64, 74, 52, 51, 65],
      },
      {
        label: 'Paid3',
        fill: false,
        backgroundColor: '#666666',
        borderColor: '#666666',
        data: [24, 1, 2, 3, 4, 5, 6],
      },
    ],
  },
  options: {
    responsive: true,
    /**
     * Default legends are ugly and impossible to style.
     * See examples in charts.html to add your own legends
     *  */
    legend: {
      display: false,
    },
    tooltips: {
      mode: 'index',
      intersect: false,
    },
    hover: {
      mode: 'nearest',
      intersect: true,
    },
    scales: {
      x: {
        display: true,
        scaleLabel: {
          display: true,
          labelString: 'Round',
        },
      },
      y: {
        display: true,
        scaleLabel: {
          display: true,
          labelString: 'Money',
        },
      },
    },
  },
}

function flushCharts() {
  const donated_history = document.getElementById('donated_history')
  window.myLine = new Chart(donated_history, lineConfig)
}

setInterval(flushCharts,500)

const money_history = document.getElementById('money_history')
window.myLine2 = new Chart(money_history, lineConfig)