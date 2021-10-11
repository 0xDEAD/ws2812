package ws2812

import (
  "machine"
  ws "tinygo.org/x/drivers/ws2812"
)

type Stripe struct {
  device ws.Device
}

func New(length int, pin machine.Pin) {
  return Stripe{
    device = ws.New(pin),
  }
}
