pub const position = struct {
    x: u32,
    y: i32,
};
pub var default_position = position{
    .x = 50,
    .y = -34,
};
pub const player = struct {
    pos: position,
    texture: []u8,
};
pub var p = player{
    .pos = default_position,
    .texture = "res/image.png",
};
