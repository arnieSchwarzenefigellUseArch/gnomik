document.addEventListener('DOMContentLoaded', function() {
    const numberInput = document.getElementById('number');
    const convertBtn = document.getElementById('convert');
    const resultDiv = document.getElementById('result');
    const errorDiv = document.getElementById('error');
    const decimalSpan = document.getElementById('decimal');
    const binarySpan = document.getElementById('binary');
    
    convertBtn.addEventListener('click', convertNumber);
    
    numberInput.addEventListener('keypress', function(e) {
        if (e.key === 'Enter') {
            convertNumber();
        }
    });
    
    async function convertNumber() {
        const number = parseInt(numberInput.value);
        
        if (isNaN(number) || number < 0) {
            showError('Введите положительное число');
            return;
        }
        
        try {
            // ОТПРАВЛЯЕМ ЗАПРОС НА СЕРВЕР
            const response = await fetch('/api/convert', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ number: number })
            });
            
            const data = await response.json();
            
            if (data.error) {
                showError(data.error);
                return;
            }
            
            // ПОКАЗЫВАЕМ РЕЗУЛЬТАТ
            showResult(data.decimal, data.binary);
            
        } catch (error) {
            showError('Ошибка соединения');
        }
    }
    
    function showResult(decimal, binary) {
        errorDiv.style.display = 'none';
        decimalSpan.textContent = decimal;
        binarySpan.textContent = binary;
        resultDiv.style.display = 'block';
    }
    
    function showError(message) {
        resultDiv.style.display = 'none';
        errorDiv.textContent = message;
        errorDiv.style.display = 'block';
    }
});