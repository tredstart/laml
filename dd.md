1) structure:
 - files of struct definitions
 - files of structs themselves
    - create a pub var for every new instance of a struct without a specific name

2) syntatic structure:
 - ":" marks start of the struct
 - ";" marks end of the struct
 - "var" marks a variable without a type
 - "#" marks an import
 - "{}" is a tuple/array

=# example #=

```components.ldef
velocity:
    i32 x,
    u8 y,
;

position: u32 x, i32 y;
player_def: 
    position pos,
    velocity vel,
    string texture,
;
box:
    u32 width,
    u32 height, 
    position pos,
;
```

```player.lcl
cmp = import "components.lcomp"
cmp.position default_pos:
    x = 50,
    y = 50,
;

cmp.player_def player:
    position = default_pos,
    velocity = {x=0, y=70},
    texture = "res/image.png",
;

cmp.box box1: width = 50, height = 50, pos = default_pos;
```

which will translate to
```components_def.zig
pub const velocity = struct {
    x i32,
    y u8,
};
pub const position = struct {
    x i32,
    y u8,
};
pub const player_def = struct {
    pos position,
    vel velocity,
    image []u8,
};
pub const box = struct {
    width u32,
    height u32, 
    pos position,

};
```

```player_obj.zig
const components_def = @import("components_def.zig");
pub var default_pos = components_def.position{.x=50, .y=50};
pub var player = components_def.player_def{.pos=default_pos, .vel=.{.x=0, .y=70,}, .texture="res/image.png"};
pub var box1 = components_def.box{.width = 50, .height = 50, .pos=default_pos};
```
