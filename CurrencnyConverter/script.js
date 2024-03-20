function convert() {
    var amount = parseFloat(document.getElementById('amount').value);
    var fromCurrency = document.getElementById('from').value;
    var toCurrency = document.getElementById('to').value;
  
    // You can replace this with an API call to get real-time exchange rates
    var exchangeRates = {
      USD: {
        EUR: 0.88,
        INR: 83.15,
      },
      EUR: {
        USD: 1.14,
        INR: 90.15
      },
      INR: {
        USD: 0.012,
        EUR: 0.011
      },
    };
  
    if (exchangeRates[fromCurrency] && exchangeRates[fromCurrency][toCurrency]) {
      var result = amount * exchangeRates[fromCurrency][toCurrency];
      document.getElementById('result').innerHTML = amount + ' ' + fromCurrency + ' = ' + result.toFixed(2) + ' ' + toCurrency;
    } else {
      document.getElementById('result').innerHTML = 'Cannot convert from ' + fromCurrency + ' to ' + toCurrency;
    }
  }
  