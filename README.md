

### Object

1. 商品
    1. ProductID: 
    2. 價格: float
2. 用戶
    1. 等級: string
    2. cash: float
    3. point: interger
3. 收銀系統
    1. 收銀機
        1. 加總
    2. Scaner(Based on Single)
        1. 輸入會員 ScanMember
        2. 掃描商品 ScanProduct
    3. Pay
    
    ```mermaid
    classDiagram
    Pay~Interface~ <|-- NormalPay
    Pay~Interface~ <|-- VipPay
    Pay~Interface~ <|-- PointDiscount
    Pay~Interface~ <|-- AnotherPay
    Cashier *-- Pay~Interface~
    Pay~Interface~ : performPay(Member)
    
    class Cashier{
    	+pay()
    	-sum()
    }
    class NormalPay{
      +pay(Member)
    }
    class VipPay{
    	+pay(Member)
    }
    class PointDiscount{
    	+pay(Member)
    }
    class AnotherPay{
    	+pay(Member)
    }
    class Product{
      -price: float
    	+SetPrice()
    }
    class member{
    	-level: str
    	-cash: float
    	-point: float
    	+SetCash()
    	+SetLevel()
    	+SetPoint()
    	+GetCash()
    	+GetPoint()
    	+GetLevel()
    }
    ```