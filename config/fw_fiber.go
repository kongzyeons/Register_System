package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Fiber_fw struct {
	engine *fiber.App
}

func NewframeworkFiber() *Fiber_fw {
	return &Fiber_fw{}
}

func (f *Fiber_fw) Default() {
	f.engine = fiber.New()
	f.engine.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
}

func (f *Fiber_fw) Get(pathurl string, api func(*fiber.Ctx) error) {
	f.engine.Get(pathurl, api)
}
func (f *Fiber_fw) Post(pathurl string, api func(*fiber.Ctx) error) {
	f.engine.Post(pathurl, api)
}
func (f *Fiber_fw) Delete(pathurl string, api func(*fiber.Ctx) error) {
	f.engine.Delete(pathurl, api)
}
func (f *Fiber_fw) Put(pathurl string, api func(*fiber.Ctx) error) {
	f.engine.Put(pathurl, api)
}

func (f *Fiber_fw) Group(pathurl string, api func(*fiber.Ctx) error) fiber.Router {
	return f.engine.Group(pathurl, api)
}

func (f *Fiber_fw) Run(portWebServie string) {
	f.engine.Listen(portWebServie) //localhost:8000
}

type FiberCtx struct {
	ctx *fiber.Ctx
}

func NewFiberCtx(c *fiber.Ctx) *FiberCtx {
	return &FiberCtx{ctx: c}
}

func (c *FiberCtx) BodyParser(v interface{}) error {
	return c.ctx.BodyParser(v)
}

func (c *FiberCtx) ParamsID(s string) string {
	return c.ctx.Params(s)
}

func (c *FiberCtx) Locals(s string) interface{} {
	return c.ctx.Locals(s)
}

func (c *FiberCtx) Status(statuscode int) *fiber.Ctx {
	return c.ctx.Status(statuscode)
}

func (c *FiberCtx) JSON(statuscode int, v interface{}) error {
	return c.ctx.Status(statuscode).JSON(v)
}
