
// A trait to define the typeclasses implementing a pricing
// functionality
pub trait Pricing {
    fn calculate(&self, units: u32) -> f32;
}

// A pricing policy for no discount
pub struct DefaultPrice {
    price_per_units: f32
}

impl Pricing for DefaultPrice {
    fn calculate(&self, units: u32) -> f32 {
        units as f32 * self.price_per_units
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
}
