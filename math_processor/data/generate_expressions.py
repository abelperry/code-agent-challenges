#!/usr/bin/env python3
import random
import math

def generate_expressions(count=1000):
    """生成1000个数学表达式"""
    expressions = []
    
    # 基础运算符
    ops = ['+', '-', '*', '/']
    
    # 函数列表
    functions = [
        'sqrt', 'pow', 'abs', 'max', 'min', 'factorial', 'fibonacci', 
        'gcd', 'lcm', 'isprime', 'sin', 'cos', 'tan', 'log', 'exp', 'ln',
        'ceil', 'floor', 'round', 'sum', 'avg', 'median', 'variance', 'stdev'
    ]
    
    # 常用角度
    angles = [0, 30, 45, 60, 90, 120, 135, 150, 180]
    
    # 常用数字
    small_nums = list(range(1, 21))
    medium_nums = [25, 30, 36, 40, 45, 50, 60, 64, 70, 75, 80, 90, 100]
    large_nums = [120, 144, 150, 180, 200, 225, 250, 300, 360, 400, 500]
    
    for i in range(count):
        expr_type = random.choice([
            'basic_arithmetic', 'single_function', 'combined_functions', 
            'complex_expression', 'statistical', 'trigonometric', 
            'logarithmic', 'factorial_fibonacci', 'number_theory'
        ])
        
        if expr_type == 'basic_arithmetic':
            # 基础四则运算
            a, b, c = random.choices(small_nums + medium_nums, k=3)
            op1, op2 = random.choices(ops, k=2)
            expressions.append(f"{a}{op1}{b}{op2}{c}")
            
        elif expr_type == 'single_function':
            # 单个函数
            func = random.choice(['sqrt', 'abs', 'ceil', 'floor', 'round'])
            if func == 'sqrt':
                # 使用完全平方数
                squares = [4, 9, 16, 25, 36, 49, 64, 81, 100, 121, 144, 169, 196, 225]
                num = random.choice(squares)
                expressions.append(f"sqrt({num})")
            elif func == 'abs':
                num = random.choice(range(-100, -1))
                expressions.append(f"abs({num})")
            else:
                num = round(random.uniform(1, 20), 1)
                expressions.append(f"{func}({num})")
                
        elif expr_type == 'combined_functions':
            # 组合函数
            func1 = random.choice(['sqrt', 'pow', 'abs'])
            func2 = random.choice(['sqrt', 'pow', 'abs'])
            op = random.choice(ops)
            
            if func1 == 'pow':
                base, exp = random.choice([(2,3), (2,4), (3,2), (3,3), (4,2), (5,2)])
                expr1 = f"pow({base},{exp})"
            elif func1 == 'sqrt':
                num = random.choice([4, 9, 16, 25, 36, 49, 64, 81, 100])
                expr1 = f"sqrt({num})"
            else:
                num = random.choice(range(-50, -1))
                expr1 = f"abs({num})"
                
            if func2 == 'pow':
                base, exp = random.choice([(2,2), (2,3), (3,2), (4,2), (5,2)])
                expr2 = f"pow({base},{exp})"
            elif func2 == 'sqrt':
                num = random.choice([4, 9, 16, 25, 36, 49, 64, 81, 100])
                expr2 = f"sqrt({num})"
            else:
                num = random.choice(range(-30, -1))
                expr2 = f"abs({num})"
                
            expressions.append(f"{expr1}{op}{expr2}")
            
        elif expr_type == 'complex_expression':
            # 复杂表达式
            patterns = [
                lambda: f"sqrt({random.choice([16, 25, 36, 49])}) + pow({random.choice([2,3])},{random.choice([2,3])})",
                lambda: f"abs({random.choice(range(-20, -1))}) * sqrt({random.choice([4, 9, 16, 25])})",
                lambda: f"max({random.choice(small_nums)},{random.choice(small_nums)},{random.choice(small_nums)}) - min({random.choice(small_nums)},{random.choice(small_nums)})",
                lambda: f"pow({random.choice([2,3,4])},{random.choice([2,3])}) + {random.choice(small_nums)} * {random.choice(small_nums)}",
                lambda: f"sqrt({random.choice([36, 49, 64, 81, 100])}) / sqrt({random.choice([4, 9, 16, 25])})"
            ]
            expressions.append(random.choice(patterns)())
            
        elif expr_type == 'statistical':
            # 统计函数
            func = random.choice(['sum', 'avg', 'median'])
            nums = random.choices(range(1, 21), k=random.randint(3, 6))
            nums_str = ','.join(map(str, nums))
            expressions.append(f"{func}({nums_str})")
            
        elif expr_type == 'trigonometric':
            # 三角函数
            func = random.choice(['sin', 'cos', 'tan'])
            angle = random.choice(angles)
            
            # 有时组合三角函数
            if random.random() < 0.3:
                func2 = random.choice(['sin', 'cos', 'tan'])
                angle2 = random.choice(angles)
                op = random.choice(['+', '-', '*'])
                expressions.append(f"{func}({angle}){op}{func2}({angle2})")
            else:
                expressions.append(f"{func}({angle})")
                
        elif expr_type == 'logarithmic':
            # 对数指数函数
            patterns = [
                lambda: f"log({random.choice([10, 100, 1000])})",
                lambda: f"exp({random.choice([0, 1, 2, 3])})",
                lambda: f"ln({random.choice([1, 2, 3, 4, 5, 10])})",
                lambda: f"log({random.choice([100, 1000])}) / log({random.choice([10])})",
                lambda: f"exp({random.choice([1, 2])}) + exp({random.choice([0, 1])})"
            ]
            expressions.append(random.choice(patterns)())
            
        elif expr_type == 'factorial_fibonacci':
            # 阶乘和斐波那契
            if random.random() < 0.5:
                n = random.choice(range(3, 9))
                if random.random() < 0.3:
                    expressions.append(f"{n}!")
                else:
                    expressions.append(f"factorial({n})")
            else:
                n = random.choice(range(5, 16))
                expressions.append(f"fibonacci({n})")
                
        elif expr_type == 'number_theory':
            # 数论函数
            func = random.choice(['gcd', 'lcm', 'isprime'])
            if func in ['gcd', 'lcm']:
                # 选择有意义的数对
                pairs = [(12,18), (24,36), (15,25), (20,30), (8,12), (6,9), (14,21), (16,24)]
                a, b = random.choice(pairs)
                expressions.append(f"{func}({a},{b})")
            else:
                # isprime
                primes = [11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47]
                composites = [12, 14, 15, 16, 18, 20, 21, 22, 24, 25, 26, 27, 28]
                num = random.choice(primes + composites)
                expressions.append(f"isprime({num})")
    
    return expressions

# 生成表达式
new_expressions = generate_expressions(100000)

# 追加到文件
with open('problems.txt', 'a') as f:
    for expr in new_expressions:
        f.write(f"{expr}\n")

print(f"Successfully generated and appended {len(new_expressions)} expressions to problems.txt")
