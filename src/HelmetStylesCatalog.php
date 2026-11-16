<?php

declare(strict_types=1);

namespace CustomGraphics\HelmetStyles;

use RuntimeException;

final class HelmetStylesCatalog
{
    /** @var array{updated: string, zones: list<string>, finishes: list<string>, styles: list<array{id: string, category: string, model: string, hasPeak: bool, hasChinBar: bool}>}|null */
    private static ?array $cache = null;

    /**
     * @return array{updated: string, zones: list<string>, finishes: list<string>, styles: list<array{id: string, category: string, model: string, hasPeak: bool, hasChinBar: bool}>}
     */
    public static function raw(): array
    {
        if (self::$cache !== null) {
            return self::$cache;
        }
        $candidates = [
            dirname(__DIR__) . '/styles.json',
            dirname(__DIR__) . '/data/styles.json',
        ];
        foreach ($candidates as $path) {
            if (is_file($path)) {
                $json = file_get_contents($path);
                if ($json === false) {
                    throw new RuntimeException('Unable to read styles.json');
                }
                /** @var array{updated: string, zones: list<string>, finishes: list<string>, styles: list<array{id: string, category: string, model: string, hasPeak: bool, hasChinBar: bool}>} $decoded */
                $decoded = json_decode($json, true, 512, JSON_THROW_ON_ERROR);
                self::$cache = $decoded;
                return $decoded;
            }
        }
        throw new RuntimeException('styles.json dataset not found');
    }

    /**
     * @return list<array{id: string, category: string, model: string, hasPeak: bool, hasChinBar: bool}>
     */
    public static function styles(): array
    {
        return self::raw()['styles'];
    }

    /**
     * @return list<string>
     */
    public static function finishes(): array
    {
        return self::raw()['finishes'];
    }

    /**
     * @return list<string>
     */
    public static function zones(): array
    {
        return self::raw()['zones'];
    }

    /**
     * Find a helmet shell style by id or model name (case and separator insensitive).
     *
     * @return array{id: string, category: string, model: string, hasPeak: bool, hasChinBar: bool}|null
     */
    public static function findStyle(string $query): ?array
    {
        $q = self::normalize($query);
        if ($q === '') {
            return null;
        }
        foreach (self::styles() as $style) {
            if (self::normalize($style['id']) === $q || self::normalize($style['model']) === $q) {
                return $style;
            }
        }
        foreach (self::styles() as $style) {
            if (str_contains(self::normalize($style['model']), $q) || str_contains(self::normalize($style['id']), $q)) {
                return $style;
            }
        }
        return null;
    }

    /**
     * Return all helmet styles in a rider category (Full face, Motocross, Adventure, Open face, Half helmet).
     *
     * @return list<array{id: string, category: string, model: string, hasPeak: bool, hasChinBar: bool}>
     */
    public static function byCategory(string $category): array
    {
        $norm = self::normalize($category);
        $matched = [];
        foreach (self::styles() as $style) {
            if (self::normalize($style['category']) === $norm) {
                $matched[] = $style;
            }
        }
        return $matched;
    }

    /**
     * Return paintable shell zones available on a given helmet style.
     *
     * @return list<string>
     */
    public static function paintableZones(string $query): array
    {
        $style = self::findStyle($query);
        if ($style === null) {
            return [];
        }
        if ($style['hasChinBar']) {
            return self::zones();
        }
        return array_values(array_filter(
            self::zones(),
            static fn (string $zone): bool => $zone !== 'chin'
        ));
    }

    private static function normalize(string $value): string
    {
        return (string) preg_replace('/[^a-z0-9]/', '', strtolower($value));
    }
}
