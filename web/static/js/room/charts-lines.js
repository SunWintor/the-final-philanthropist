
historyLine = null
moneyLine = null

function initDatasets(map) {
  datasets = []
  $.each(map,function(key,value) {
    datasets.push(value)
  });
  return datasets
}

function initLine() {
  if (Object.keys(playerHistoryMap).length > 0 && historyLine == null) {
    historyConfig = {
      type: 'line',
      data: {
        labels: [],
        datasets: initDatasets(playerHistoryMap),
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
    historyLine = new Chart(document.getElementById('donated_history'), historyConfig)
  }
  if (Object.keys(playerMoneyMap).length > 0 && moneyLine == null) {
    moneyConfig = {
      type: 'line',
      data: {
        labels: [],
        datasets: initDatasets(playerMoneyMap),
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
    moneyLine = new Chart(document.getElementById('money_history'), moneyConfig)
  }
}

function flushCharts() {
  initLine()
  $.each(playerHistoryMap,function(key,value){
    addData(historyLine, value, value.index)
  });
  $.each(playerMoneyMap,function(key,value){
    addData(moneyLine, value, value.index)
  });
}

function addData(chart, dataset, datasetIndex) {
  chart.data.datasets[datasetIndex].data = dataset.data;
  chart.data.labels = chartsLabels;
  chart.update();
}

setInterval(flushCharts,3000)