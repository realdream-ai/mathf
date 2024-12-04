package mathf

import (
	"godot-ext/mathf/impl"
)

type Rect2i struct {
	Position Vector2i
	Size     Vector2i
}

func (v Rect2i) toImpl() impl.Rect2i {
	return impl.Rect2i{
		Position: v.Position.toImpl(),
		Size:     v.Size.toImpl(),
	}
}

func (v Rect2i) fromImpl(iv impl.Rect2i) Rect2i {
	return Rect2i{
		Position: NewVector2i(int(iv.Position[0]), int(iv.Position[1])),
		Size:     NewVector2i(int(iv.Size[0]), int(iv.Size[1])),
	}
}

func (v *Rect2i) toImplf() impl.Rect2 {
	return impl.Rect2{
		Position: v.Position.toImplf(),
		Size:     v.Size.toImplf(),
	}
}

func (v *Rect2i) fromImplf(iv impl.Rect2) Rect2i {
	return Rect2i{
		Position: NewVector2i(int(iv.Position[0]), int(iv.Position[1])),
		Size:     NewVector2i(int(iv.Size[0]), int(iv.Size[1])),
	}
}

func NewRect2i(x, y, w, h int) Rect2i {
	return Rect2i{
		Position: NewVector2i(x, y),
		Size:     NewVector2i(w, h),
	}
}

func (r *Rect2i) End() Vector2i {
	impl := r.toImpl()
	end := impl.End()
	return NewVector2i(int(end[0]), int(end[1]))
}

func (r *Rect2i) SetEnd(end Vector2i) {
	impl := r.toImpl()
	impl.SetEnd(end.toImpl())
	*r = r.fromImpl(impl)
}

func (r *Rect2i) Rect2() Rect2 {
	return NewRect2(float64(r.Position.X), float64(r.Position.Y), float64(r.Size.X), float64(r.Size.Y))
}

func (r *Rect2i) Abs() Rect2i {
	return r.fromImpl(r.toImpl().Abs())
}

func (r *Rect2i) Encloses(b Rect2i) bool {
	return r.toImpl().Encloses(b.toImpl())
}

func (r *Rect2i) Expand(to Vector2i) Rect2i {
	return r.fromImpl(r.toImpl().Expand(to.toImpl()))
}

func (r *Rect2i) Area() int64 {
	return r.toImpl().Area()
}

func (r *Rect2i) Center() Vector2i {
	impl := r.toImpl()
	center := impl.Center()
	return NewVector2i(int(center[0]), int(center[1]))
}

func (r *Rect2i) Grow(amount int64) Rect2i {
	return r.fromImpl(r.toImpl().Grow(amount))
}

func (r *Rect2i) GrowIndividual(left, top, right, bottom int64) Rect2i {
	return r.fromImpl(r.toImpl().GrowIndividual(left, top, right, bottom))
}

func (r *Rect2i) GrowSide(side Side, amount int64) Rect2i {
	return r.fromImpl(r.toImpl().GrowSide(side, amount))
}

func (r *Rect2i) HasArea() bool {
	return r.toImpl().HasArea()
}

func (r *Rect2i) HasPoint(point Vector2i) bool {
	return r.toImpl().HasPoint(point.toImpl())
}

func (r *Rect2i) Intersection(b Rect2i) Rect2i {
	return r.fromImpl(r.toImpl().Intersection(b.toImpl()))
}

func (r *Rect2i) Intersects(b Rect2i, includeBorders bool) bool {
	return r.toImpl().Intersects(b.toImpl(), includeBorders)
}

func (r *Rect2i) Merge(b Rect2i) Rect2i {
	return r.fromImpl(r.toImpl().Merge(b.toImpl()))
}