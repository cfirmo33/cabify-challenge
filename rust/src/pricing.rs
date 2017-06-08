
// A trait to define the typeclasses implementing a pricing
// functionality
pub trait Pricing : Send {
    fn calculate(&self, units: u32) -> f32;
}

// A pricing policy for no discount
pub struct DefaultPrice {
    pub price_per_units: f32
}

impl Pricing for DefaultPrice {
    fn calculate(&self, units: u32) -> f32 {
        units as f32 * self.price_per_units
    }
}

// A pricing policy get two and pay one discount.
pub struct GetTwoPayOne {
    pub price_per_units: f32
}

impl Pricing for GetTwoPayOne {
    fn calculate(&self, units: u32) -> f32 {
        (units - units / 2) as f32 * self.price_per_units
    }
}
// A pricing policy bulk discount
pub struct BulkDiscount {
    pub price_per_units: f32,
    pub discount_price_per_units: f32,
    pub min_bulk_units: u32,
}

impl Pricing for BulkDiscount {
    fn calculate(&self, units: u32) -> f32 {
        if units < self.min_bulk_units {
            units as f32 * self.price_per_units
        } else {
            units as f32 * self.discount_price_per_units
        }
    }
}


#[cfg(test)]
mod test {

    use super::*;

    #[test]
    fn test_default_price() {
        let p = DefaultPrice { price_per_units: 1.95 };

        assert_eq!(0.00, p.calculate(0));
        assert_eq!(1.95, p.calculate(1));
        assert_eq!(3.90, p.calculate(2));
    }

    #[test]
    fn test_get_two_pay_one() {
        let p = GetTwoPayOne { price_per_units: 1.95 };

        assert_eq!(0.00, p.calculate(0));
        assert_eq!(1.95, p.calculate(1));
        assert_eq!(1.95, p.calculate(2));
        assert_eq!(3.90, p.calculate(3));
        assert_eq!(3.90, p.calculate(4));
    }

    #[test]
    fn test_bulk_discount() {
        let p = BulkDiscount {
            price_per_units: 1.95,
            discount_price_per_units: 1.75,
            min_bulk_units: 3
        };

        assert_eq!(0.00, p.calculate(0));
        assert_eq!(1.95, p.calculate(1));
        assert_eq!(3.90, p.calculate(2));
        assert_eq!(5.25, p.calculate(3));
        assert_eq!(7.00, p.calculate(4));
    }
}
