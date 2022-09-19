
historyLine = null
moneyLine = null

function initDatasets(map) {
  datasets = []
  $.each(map,function(key,value) {
    datasets.push(value)
  });
  return datasets
}

function initExampleTemplate() {
  exp = ""
  $.each(playerHistoryChartList,function(key, value) {
    exp += lineExampleTemplate.signMix(value.backgroundColor, value.label)
  });
  $('#donated_history_exp').html(exp)
  exp2 = ""
  $.each(playerMoneyChartList,function(key, value) {
    exp2 += lineExampleTemplate.signMix(value.backgroundColor, value.label)
  });
  $('#money_history_exp').html(exp2)
}

function initLine() {
  if (Object.keys(playerHistoryChartList).length > 0 && historyLine == null) {
    initExampleTemplate()
    historyConfig = {
      type: 'line',
      data: {
        labels: [],
        datasets: initDatasets(playerHistoryChartList),
      },
      options: {
        elements: {
          line: {
            tension: 0 // 禁用贝塞尔曲线
          }
        },
        animation: {
          duration: 0 // 一般动画时间
        },
        responsiveAnimationDuration: 0, // 调整大小后的动画持续时间
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
          animationDuration: 0 // 悬停项目时动画的持续时间
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
  if (Object.keys(playerMoneyChartList).length > 0 && moneyLine == null) {
    moneyConfig = {
      type: 'line',
      data: {
        labels: [],
        datasets: initDatasets(playerMoneyChartList),
      },
      options: {
        elements: {
          line: {
            tension: 0 // 禁用贝塞尔曲线
          }
        },
        animation: {
          duration: 0 // 一般动画时间
        },
        responsiveAnimationDuration: 0, // 调整大小后的动画持续时间
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
          animationDuration: 0 // 悬停项目时动画的持续时间
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
  $.each(playerHistoryChartList,function(key, value){
    addData(historyLine, value, value.index)
  });
  $.each(playerMoneyChartList,function(key, value){
    addData(moneyLine, value, value.index)
  });
}

function addData(chart, dataset, datasetIndex) {
  chart.data.datasets[datasetIndex].data = dataset.data;
  chartsLabels = []
  index = 0
  for (let player of dataset.data) {
    chartsLabels.push(index++)
  }
  chart.data.labels = chartsLabels;
  chart.update();
}

function clearTfpChara() {
  clearChart(historyLine)
  clearChart(moneyLine)
  historyLine = null
  moneyLine = null
}

function clearChart(chart) {
  if (!chart) {
    return
  }
  chart.data.labels.pop();
  chart.data.datasets.forEach(dataset => {
    dataset.data.pop();
  });
  chart.update();
  initExampleTemplate()
}

setInterval(flushCharts,3000)