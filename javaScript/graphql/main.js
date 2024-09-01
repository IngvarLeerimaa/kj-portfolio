// Test ratio = how many test have user passed / how many test cases were run 0 flase 1 true
// (remember to omit null and float grades(floats == audits))
// Basic user identification firstname lastname email login
// uses normal, nested and arguments queries
const graphqlQuery = `
    query{
        user { 
            firstName
            lastName
            email
            login
            progresses{
                grade
              }
            transactions(where: { type: { _eq: "xp" } }, order_by: { createdAt: asc }) {
                id
                type
                amount
                path
                createdAt
            }
        }
    }
          `;

async function query(){
    const token = localStorage.getItem('token');
    if (!token){
        location.replace('index.html');
    }

    try {
        const response = await fetch('https://01.kood.tech/api/graphql-engine/v1/graphql',{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`
            },
            body: JSON.stringify({query: graphqlQuery})
        });
        const data = await response.json(); // Read the body

        // Extract from object
        const userData = data.data.user[0];

        const user = {
            login: userData.login,
            firstName: userData.firstName,
            lastName: userData.lastName,
            email: userData.email
        };

        // For test ratio during diffrent piscines
        const progresses = userData.progresses;
        let passedRatio = calcPassedRatio(progresses);

        // For xp amount overtime
        const rawTransactions = userData.transactions;
        const transactions = rawTransactions.filter(transaction => !transaction.path.includes('piscine'));

        // Render the query elements
        renderUser(user);
        renderCircleChart(passedRatio);
        const xp = cleanXp(transactions);
        document.getElementById('basicXp').textContent = `You have earned ${xp} xp during your time at kood/jõhvi!`;
        renderXp(transactions);
        addLogOutButton();
       }
    catch (error){
        console.error('Error:', error);
    }

}
//inits script
query();
// Calculates the ratio of tests passed and returns test passed, percentage and total test taken
function calcPassedRatio(progresses){
    let passed = 0;
    let total = 0;
    for (const progress of progresses){
        if (progress.grade !== null){
            if (progress.grade == 1){
                passed++;
            }

            if (progress.grade == 0 || progress.grade == 1) {
                total++;
            }
        }
    }
    let passedRatio = passed / total;
    let percentage = (passedRatio * 100).toFixed(0);
    return { passed: passed, percentage: percentage, total: total };
}

// Adds up all xp
function cleanXp(transactions){
    let xp = 0;
    for (const transaction of transactions){
        xp += transaction.amount;
    }
    return xp;
}

function renderUser(user){
    // Get the user div
    const userDiv = document.getElementById('user');

    // Clear the userDiv
    userDiv.innerHTML = '';

    // Create h2 element for login
    const h2 = document.createElement('h2');
    h2.textContent = user.login;
    userDiv.appendChild(h2);

    // Create p element for name
    const pName = document.createElement('p');
    pName.textContent = `Name: ${user.firstName} ${user.lastName}`;
    userDiv.appendChild(pName);

    // Create p element for email
    const pEmail = document.createElement('p');
    pEmail.textContent = `Email: ${user.email}`;
    userDiv.appendChild(pEmail);

    // Create p element for xp and set it later
    const xp = document.createElement('p');
    xp.id = 'basicXp';
    xp.textContent = ``;
    userDiv.appendChild(xp);


}

// Render the circle chart
function renderCircleChart(passedRatio){
    const sideTextTitle = document.getElementById('better');
    sideTextTitle.textContent = `You have passed ${passedRatio.percentage}% of the tests during your time at kood/jõhvi!`;
    const sideTextP = document.getElementById('sideTextP');
    sideTextP.textContent = `Remember all the test you tried during different pisicines? Out of ${passedRatio.total} tests you passed ${passedRatio.passed} of them!`;
    
    //Pie chart animation    
    const chart = document.querySelector('[data-pie]');
    let speed = +chart.dataset.speed || 0;
    const circle = chart.querySelector("circle");
    const text = chart.querySelector("text");
    text.textContent = `${passedRatio.percentage}%`;
    chart.setAttribute("aria-label", `${passedRatio.percentage} percent pie chart`);
    circle.animate(
      [
        {
          strokeDashoffset: 100,
        },
        {
          strokeDashoffset: 100 - passedRatio.percentage,
        },
      ],
      {
        duration: speed,
        easing: "cubic-bezier(0.57,-0.04, 0.41, 1.13)",
        fill: "forwards",
      }
    );
    text.animate(
      [
        {
          opacity: 0,
          transform: "translateY(20%)",
        },
        {
          opacity: 1,
          transform: "translateY(0%)",
        },
      ],
      {
        easing: "cubic-bezier(0.57,-0.04, 0.41, 1.13)",
        fill: "forwards",
      }
    );



}

// Render the bar chart
function renderXp(transactions) {
    const chartContainer = document.getElementById('chart-container');
    const tooltip = document.getElementById('tooltip');
    const svgns = 'http://www.w3.org/2000/svg';
    const svg = document.createElementNS(svgns, 'svg');
    svg.setAttribute('width', '100%');
    svg.setAttribute('height', '100%');
  
    const taskMap = {};
  
    // Extract task names and sum XP for each task
    transactions.forEach(transaction => {
      const task = transaction.path.replace('/johvi/div-01/', '');
      if (!taskMap[task]) {
        taskMap[task] = [];
      }
      taskMap[task].push({ amount: transaction.amount, date: new Date(transaction.createdAt) });
    });
  
    const tasks = Object.keys(taskMap);
    const maxXP = Math.max(...tasks.map(task => taskMap[task].reduce((acc, curr) => acc + curr.amount, 0)));
    const barWidth = (700 - (tasks.length - 1) * 3) / tasks.length; //based on the width
    const barHeightRatio = 350 / maxXP; //based on height 
  
    tasks.forEach((task, index) => {
      const x = index * (barWidth + 6); 
      const y = 400 - taskMap[task].reduce((acc, curr) => acc + curr.amount, 0) * barHeightRatio;
      const height = taskMap[task].reduce((acc, curr) => acc + curr.amount, 0) * barHeightRatio;
  
      // Create bar
      const bar = document.createElementNS(svgns, 'rect');
      bar.setAttribute('class', 'bar');
      bar.setAttribute('x', x);
      bar.setAttribute('y', y);
      bar.setAttribute('width', barWidth);
      bar.setAttribute('height', height);
  
      // Add event listener for tooltip
      bar.addEventListener('mouseover', function(event) {
        const totalAmount = taskMap[task].reduce((acc, curr) => acc + curr.amount, 0);
        const dates = taskMap[task].map(item => item.date.toLocaleDateString()).join(', ');
        const tooltipText = `You completed ${task} on ${dates} and recived ${totalAmount} XP.`;
        tooltip.textContent = tooltipText;
        tooltip.style.left = (event.pageX + 10) + 'px'; //pos adjust 
        tooltip.style.top = (event.pageY - 20) + 'px'; //pos adjust
        tooltip.style.display = 'block';
      });
  
      bar.addEventListener('mouseout', function() {
        tooltip.style.display = 'none';
      });
  
      svg.appendChild(bar);
    });
  
    chartContainer.appendChild(svg);
  }
function addLogOutButton(){
    const logOutButton = document.createElement('h2');
    logOutButton.style.textAlign = 'center';
    logOutButton.textContent = 'Log Out';
    logOutButton.style.cursor = 'pointer';
    logOutButton.addEventListener('click', function(){
        localStorage.removeItem('token');
        location.replace('index.html');
    });
    document.querySelector('main').appendChild(logOutButton);
}

