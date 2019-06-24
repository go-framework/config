package redis

// Register update function.
func (c *Config) UpdateFuncRegister(f func(config *Config)) {
	c.updateFunc = append(c.updateFunc, f)
}
