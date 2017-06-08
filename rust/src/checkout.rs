use std::collections::HashMap;
use std::sync::mpsc;
use std::thread;

use pricing::*;

pub type PricingRules = HashMap<String, Box<Pricing>>;

pub struct Checkout {
    channel: mpsc::Sender<CheckoutMessage>
}

impl Checkout {
    pub fn new(rules: PricingRules) -> Checkout {
        let (tx, rx) = mpsc::channel();
        thread::spawn(move || {
            handle_requests(rules, rx)
        });
        Checkout { channel: tx }
    }

    pub fn scan(&mut self, product: String) {
        let msg = CheckoutMessage::Scan { product_name: product };
        self.channel.send(msg).unwrap()
    }

    pub fn total(&self) -> f32 {
        let (tx, rx) = mpsc::channel();
        let msg = CheckoutMessage::Total { response_channel: tx };
        self.channel.send(msg).unwrap();
        rx.recv().unwrap()
    }
}

enum CheckoutMessage {
    Scan { product_name: String },
    Total { response_channel: mpsc::Sender<f32> }
}

fn handle_requests(rules: PricingRules,
                   channel: mpsc::Receiver<CheckoutMessage>) {
    let mut cart: HashMap<String, u32> = HashMap::new();
    loop {
        let msg = channel.recv().expect("message received");
        match msg {
            CheckoutMessage::Scan { product_name } => {
                let count = cart.get(&product_name).unwrap_or(&0) + 1;
                cart.insert(product_name, count);
            },
            CheckoutMessage::Total { response_channel } => {
                let total = cart.iter().fold(0.0, |accum, (k, v)| {
                    accum + rules.get(k).unwrap().calculate(*v)
                });
                response_channel.send(total).unwrap();
                break
            },
        }
    }
}

pub fn default_pricing_rules() -> PricingRules {
    let mut rules: PricingRules = HashMap::new();
    rules.insert("VOUCHER".to_string(),
                 Box::new(GetTwoPayOne { price_per_units: 5.00 }));
    rules.insert("TSHIRT".to_string(),
                 Box::new(BulkDiscount {
                     price_per_units: 20.00,
                     discount_price_per_units: 19.00,
                     min_bulk_units: 3
                 }));
    rules.insert("MUG".to_string(),
                 Box::new(DefaultPrice { price_per_units: 7.50 }));
    rules
}

#[cfg(test)]
mod test {

    use super::*;

    #[test]
    fn test_scenario_1() {
        let mut co = Checkout::new(default_pricing_rules());
        co.scan("VOUCHER".to_string());
        co.scan("TSHIRT".to_string());
        co.scan("MUG".to_string());

        assert_eq!(32.50, co.total());
    }

    #[test]
    fn test_scenario_2() {
        let mut co = Checkout::new(default_pricing_rules());
        co.scan("VOUCHER".to_string());
        co.scan("TSHIRT".to_string());
        co.scan("VOUCHER".to_string());

        assert_eq!(25.00, co.total());
    }

    #[test]
    fn test_scenario_3() {
        let mut co = Checkout::new(default_pricing_rules());
        co.scan("TSHIRT".to_string());
        co.scan("TSHIRT".to_string());
        co.scan("TSHIRT".to_string());
        co.scan("VOUCHER".to_string());
        co.scan("TSHIRT".to_string());

        assert_eq!(81.00, co.total());
    }

    #[test]
    fn test_scenario_4() {
        let mut co = Checkout::new(default_pricing_rules());
        co.scan("VOUCHER".to_string());
        co.scan("TSHIRT".to_string());
        co.scan("VOUCHER".to_string());
        co.scan("VOUCHER".to_string());
        co.scan("MUG".to_string());
        co.scan("TSHIRT".to_string());
        co.scan("TSHIRT".to_string());

        assert_eq!(74.50, co.total());
    }
}
